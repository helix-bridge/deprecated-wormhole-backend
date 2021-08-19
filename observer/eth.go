package observer

import (
	"fmt"
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/log"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type EthTransaction struct {
	Last	int64                     `json:"last"`
	Address string                    `json:"address"`
	Method  []string                  `json:"method"`
	Result  *parallel.EtherscanResult `json:"result"`
	ch chan interface{}
}

func (e *EthTransaction) RelyOn() bool {
    return VerifyProof == e.Result.Topics[0]
}

func (e *EthTransaction) LoadData(o Observable, isRely bool) {
    needSync := func() bool {
	return isRely == e.RelyOn()
    }

    ethfrom := util.StringToInt64(string(util.HgetCache("restart", "ethfrom")))
    ethto := util.StringToInt64(string(util.HgetCache("restart", "ethto")))

    key := strings.Join(e.Method, ":")
    log.Info("eth start to load init data", "key", key, "isrely", isRely, "ethfrom", ethfrom, "ethto", ethto)
    for {
	if ethfrom >= ethto {
	    break
	}
	if eventLog, _ := parallel.EtherscanLog(ethfrom, ethfrom + 102400, e.Address, e.Method...); eventLog != nil {
	    for _, result := range eventLog.Result {
		e.Result = &result
		ethfrom = util.U256(result.BlockNumber).Int64()
		if ethfrom > ethto {
		    break
		}
		if !needSync() {
		    continue
		}
		_ = o.notify(e)
	    }
	}
	ethfrom = ethfrom + 102400
	time.Sleep(1 * time.Second)
    }
    log.Info("finish to load data", "key", key, "isrely", isRely, "ethfrom", ethfrom, "ethto", ethto)
    e.Last = ethto
}

func (e *EthTransaction) Do(o Observable) error {
	if e.Result == nil || !util.StringInSlice(e.Result.Topics[0], EthAvailableEvent) {
		return fmt.Errorf("empty transaction %s", e.Result)
	}
	return e.Redeem()
}


func (e *EthTransaction) pullEvents(o Observable) {
    old_start := e.Last
    if eventLog, _ := parallel.EtherscanLog(e.Last+1, 0, e.Address, e.Method...); eventLog != nil {
	for _, result := range eventLog.Result {
	    e.Last = util.U256(result.BlockNumber).Int64()
	    e.Result = &result
	    _ = o.notify(e)
	}
    }
    if old_start != e.Last {
	key := strings.Join(e.Method, ":")
	log.Info("set ethcan new last", "key", key, "last", e.Last)
	_ = util.SetCache(key, e.Last, 86400*7)
    }
}

func (e *EthTransaction) Pause() {
    e.ch <- true
}

func (e *EthTransaction) Resume() {
    e.ch <- false
}

func (e *EthTransaction) ErrorBreak(err error) {
    e.ch <- err
}

func (e *EthTransaction) Listen(o Observable) error {
    e.ch = make(chan interface{})
    key := strings.Join(e.Method, ":")
    if e.Last == 0 {
	if b := util.GetCache(key); b != nil {
	    e.Last = util.StringToInt64(string(b))
	} else {
	    e.Last = 8028174
	}
    }
    updateInterval := time.Second * time.Duration(10)
    updateTimer := time.NewTimer(updateInterval)
    pause := false
    go func() {
	for {
	    select {
	    case v := <-e.ch:
		switch v:= v.(type) {
		case error:
		    log.Info("observer has error", "err", v)
		    break;
		case bool:
		    if v {
			pause = true
			log.Info("ethscan event paused", "key", key, "last", e.Last)
		    } else {
			pause = false
			log.Info("ethscan event resumed", "key", key, "last", e.Last)
		    }
		}
	    case <-updateTimer.C:
		if !pause {
		    e.pullEvents(o)
		}
		updateTimer.Reset(updateInterval)
	    }
	}
    }()
    return nil
}

// https://github.com/darwinia-network/dj
func (e *EthTransaction) Redeem() error {
	logSlice := util.LogAnalysis(e.Result.Data)

	switch e.Result.Topics[0] {
	case BurnAndRedeem:
		currency := "ring"
		token := util.AddHex(e.Result.Topics[1][len(e.Result.Topics[1])-40:])
		from := util.AddHex(e.Result.Topics[2][len(e.Result.Topics[2])-40:])
		amount := decimal.NewFromBigInt(util.U256(logSlice[0]), 0)
		target := logSlice[3]
		if strings.EqualFold(token, config.Link.Kton) {
			currency = "kton"
		}
		return db.AddRedeemRecord(Eth, util.AddHex(e.Result.TransactionHash), from, target, currency, amount,
			int(util.U256(e.Result.BlockNumber).Int64()), int(util.U256(e.Result.TimeStamp).Int64()), "")

	case BurnAndRedeemDeposit:
		depositId := util.U256(e.Result.Topics[1]).Int64()
		from := util.AddHex(logSlice[0][len(logSlice[0])-40:])
		month := util.U256(logSlice[1]).Int64()
		startAt := util.U256(logSlice[2]).Int64()
		amount := decimal.NewFromBigInt(util.U256(logSlice[4]), 0)
		target := logSlice[7]
		deposit := map[string]int64{"deposit_id": depositId, "month": month, "start": startAt}

		return db.AddRedeemRecord(Eth, util.AddHex(e.Result.TransactionHash), from, target, "deposit", amount,
			int(util.U256(e.Result.BlockNumber).Int64()), int(util.U256(e.Result.TimeStamp).Int64()), util.ToString(deposit))

	case VerifyProof:
		blockNum := util.U256(logSlice[0]).Uint64()
		if e.Address == config.Link.EthereumBacking {
			db.SetTokenRegistrationConfirm(blockNum, util.AddHex(e.Result.TransactionHash))
		} else if e.Address == config.Link.TokenIssuing {
			db.SetBackingLockConfirm(blockNum, util.AddHex(e.Result.TransactionHash))
		}
		return nil

	case SetRootEvent:
		index := util.U256(logSlice[2]).Uint64()
		db.SetMMRIndexBestBlockNum(index)

	case BackingLock:
		sender := util.AddHex(e.Result.Topics[1][len(e.Result.Topics[1])-40:])
		receiver := util.AddHex(logSlice[3][len(logSlice[3])-40:])
		amount := decimal.NewFromBigInt(util.U256(logSlice[2]), 0)
		source := util.AddHex(logSlice[0][len(logSlice[0])-40:])
		target := util.AddHex(logSlice[1][len(logSlice[1])-40:])
		return db.AddEthereumLockRecord(Eth, util.AddHex(e.Result.TransactionHash), source, target, sender, receiver, amount,
			int(util.U256(e.Result.BlockNumber).Int64()), int(util.U256(e.Result.TimeStamp).Int64()))
	}
	return nil
}
