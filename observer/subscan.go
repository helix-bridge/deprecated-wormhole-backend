package observer

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"strings"
	"time"
)

type SubscanEvent struct {
	EventId  string
	ModuleId string
	Result   *parallel.SubscanEvent
	Last     int64 `json:"last"`
}

func (s *SubscanEvent) Do(o Observable) error {
	return s.Process()
}

func (s *SubscanEvent) Listen(o Observable) error {
	key := s.ModuleId + ":" + s.EventId
	if s.Last == 0 {
		if b := util.GetCache(key); b != nil {
			s.Last = util.StringToInt64(string(b))
		}
	}
	go func() {
		for {
			if eventLog := parallel.SubscanEvents(s.ModuleId, s.EventId, s.Last); eventLog != nil {
				for _, result := range eventLog {
					s.Result = &result
					if result.BlockNum > s.Last {
						s.Last = result.BlockNum
					}
					_ = o.notify(s)
				}
			}
			_ = util.SetCache(key, s.Last, 86400*7)
			time.Sleep(15 * time.Second)
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
		case "TokenRegistered":
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
