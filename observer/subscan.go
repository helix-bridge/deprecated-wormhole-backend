package observer

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/log"
	"strings"
	"time"
)

type SubscanEvent struct {
	EventId  string
	ModuleId string
	Result   *parallel.SubscanEvent
	Last     int64 `json:"last"`
	ch chan interface{}
}

func (s *SubscanEvent) RelyOn() bool {
    switch s.ModuleId {
    case strings.ToLower("EthereumRelay"):
	return false
    case strings.ToLower("EthereumBacking"):
	switch s.Result.EventId {
	case "RedeemDeposit", "RedeemKton", "RedeemRing":
	    return true
	case "LockKton", "LockRing":
	    return false
	}
    case strings.ToLower("EthereumRelayAuthorities"):
	return true
    case strings.ToLower("EthereumIssuing"):
	switch s.Result.EventId {
	case "TokenRegisterFinished":
	    return false
	case "BurnToken":
	    return false
	case "RedeemErc20":
	    return true
	}
    }
    return false
}

func (s *SubscanEvent) Do(o Observable) error {
	return s.Process()
}

func (s *SubscanEvent) LoadData(o Observable, isRely bool) {
    needSync := func() bool {
	return isRely == s.RelyOn()
    }

    restartInfo := util.HgetCacheAll("restart")
    subfrom := util.StringToInt64(restartInfo["subfrom"])
    subto := util.StringToInt64(restartInfo["subto"])

    key := s.ModuleId + ":" + s.EventId
    log.Info("subscan start to load init data", "key", key, "isrely", isRely, "subfrom", subfrom, "subto", subto)
    count := 0
    for {
	if subfrom >= subto {
	    break
	}
	if eventLog := parallel.SubscanEvents(s.ModuleId, s.EventId, subfrom); eventLog != nil {
	    count += len(eventLog)
	    for _, result := range eventLog {
		s.Result = &result
		if result.BlockNum >= subfrom {
		    subfrom = result.BlockNum + 1
		}
		if result.BlockNum > subto {
		    break
		}
		//if result.BlockNum > subfrom {
		    //subfrom = result.BlockNum
		//}
		if !needSync() {
		    continue
		}
		_ = o.notify(s)
	    }
	} else {
	    break
	}
	time.Sleep(1000 * time.Millisecond)
    }
    log.Info("finish to load data", "key", key, "isrely", isRely, "subfrom", subfrom, "subto", subto, "count", count)
    s.Last = subto
}

func (s *SubscanEvent) pullEvents(o Observable) {
    old_start := s.Last
    key := s.ModuleId + ":" + s.EventId
    if eventLog := parallel.SubscanEvents(s.ModuleId, s.EventId, s.Last); eventLog != nil {
	for _, result := range eventLog {
	    s.Result = &result
	    if result.BlockNum >= s.Last {
		s.Last = result.BlockNum + 1
	    }
	    log.Info("subscan find valid event", "key", key, "event", s.Result)
	    _ = o.notify(s)
	}
    }
    if old_start != s.Last {
	log.Info("set subscan new last", "key", key, "last", s.Last)
	_ = util.SetCache(key, s.Last, 86400*7)
    }
}

func (s *SubscanEvent) Pause() {
    s.ch <- true
}

func (s *SubscanEvent) Resume() {
    s.ch <- false
}

func (s *SubscanEvent) ErrorBreak(err error) {
    s.ch <- err
}

func (s *SubscanEvent) Listen(o Observable) error {
        s.ch = make(chan interface{})
	key := s.ModuleId + ":" + s.EventId
	if s.Last == 0 {
		if b := util.GetCache(key); b != nil {
			s.Last = util.StringToInt64(string(b))
		}
	}
	log.Info("subscan start listen", "key", key, "last", s.Last)
	updateInterval := time.Second * time.Duration(15)
	updateTimer := time.NewTimer(updateInterval)
	pause := false
	go func() {
	    for {
		select {
		case v := <-s.ch:
		    switch v:= v.(type) {
		    case error:
			log.Info("observer has error", "err", v)
			break;
		    case bool:
			if v {
			    pause = true
			    log.Info("subscan event paused", "key", key, "last", s.Last)
			} else {
			    pause = false
			    log.Info("subscan event resumed", "key", key, "last", s.Last)
			}
		    }
		case <-updateTimer.C:
		    if !pause {
			s.pullEvents(o)
		    }
		    updateTimer.Reset(updateInterval)
		}
	    }
	}()
	return nil
}

type EthereumTransactionIndex struct {
	BlockHash string `json:"col1"`
	Index     int    `json:"col2"`
}

func (s *SubscanEvent) Process() error {
	switch s.ModuleId {
	case strings.ToLower("EthereumRelay"):
		db.SetRelayBestBlockNum(util.UInt64FromInterface(s.Result.Params[0].Value))
	case strings.ToLower("EthereumBacking"):
		switch s.Result.EventId {
		case "RedeemDeposit", "RedeemKton", "RedeemRing":
			for _, param := range s.Result.Params {
				if strings.EqualFold(param.Type, "EthereumTransactionIndex") {
					var t EthereumTransactionIndex
					util.UnmarshalAny(&t, param.Value)
					if fromTx := parallel.EthGetTransactionByBlockHashAndIndex(t.BlockHash, util.IntFromInterface(t.Index)); fromTx != "" {
						db.UpdateRedeem(fromTx, s.Result.ExtrinsicIndex)
					}
				}
			}
		case "LockKton", "LockRing":
			extrinsic := parallel.SubscanExtrinsic(s.Result.ExtrinsicIndex)
			_ = db.CreateDarwiniaBacking(s.Result.ExtrinsicIndex, extrinsic)
		}
	case strings.ToLower("EthereumRelayAuthorities"):
		_ = db.MMRRootSigned(s.Result.Params)
		_ = db.MMRRootSignedForTokenRegistration(s.Result.Params)
	case strings.ToLower("EthereumIssuing"):
		extrinsic := parallel.SubscanExtrinsic(s.Result.ExtrinsicIndex)
		switch s.Result.EventId {
		case "TokenRegisterFinished":
			_ = db.CreateTokenRegisterRecord(s.Result.ExtrinsicIndex, extrinsic)
		case "BurnToken":
			_ = db.CreateTokenBurnRecord(s.Result.ExtrinsicIndex, extrinsic)
		case "RedeemErc20":
			for _, param := range s.Result.Params {
				if strings.EqualFold(param.Type, "EthereumTransactionIndex") {
					var t EthereumTransactionIndex
					util.UnmarshalAny(&t, param.Value)
					if fromTx := parallel.EthGetTransactionByBlockHashAndIndex(t.BlockHash, util.IntFromInterface(t.Index)); fromTx != "" {
						db.UpdateEthereumLockRecord(fromTx, s.Result.ExtrinsicIndex)
					}
				}
			}
		}
	}
	return nil
}
