package observer

import (
	"encoding/hex"
	"fmt"
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/crypto"
	"github.com/shopspring/decimal"
	"time"
)

type EthTransaction struct {
	Last    int64                     `json:"last"`
	Address string                    `json:"address"`
	Method  []string                  `json:"method"`
	Result  *parallel.EtherscanResult `json:"result"`
}

func (e *EthTransaction) Do(o Observable) error {
	fmt.Println("find EthTransaction", e.Result)
	if e.Result == nil || !util.StringInSlice(e.Result.Topics[0], EthAvailableEvent) {
		fmt.Println(e.Result.Topics[0], EthAvailableEvent)
		return fmt.Errorf("empty transaction %s", e.Result)
	}
	return e.RingBurnRecord()
}

func (e *EthTransaction) Listen(o Observable) error {
	key := runFuncName()
	if e.Last == 0 {
		if b := util.GetCache(key); b != nil {
			e.Last = util.StringToInt64(string(b))
		} else {
			e.Last = 8028174
		}
	}
	go func() {
		for {
			if eventLog, _ := parallel.EtherscanLog(e.Last+1, e.Address, e.Method...); eventLog != nil {
				for _, result := range eventLog.Result {
					e.Last = util.U256(result.BlockNumber).Int64()
					e.Result = &result
					_ = o.notify(e)
				}
			}
			_ = util.SetCache(key, e.Last, 86400*7)
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}

func (e *EthTransaction) RingBurnRecord() error {
	logSlice := util.LogAnalysis(e.Result.Data)

	if len(logSlice) != 4 || len(e.Result.Topics) != 3 {
		return fmt.Errorf("error log or topic %s", e.Result.TransactionHash)
	}

	address := util.AddHex(e.Result.Topics[2][len(e.Result.Topics[2])-40:])
	amount := decimal.NewFromBigInt(util.U256(logSlice[0]), 0)
	target := logSlice[3]

	currency := ring
	if e.Result.Topics[0] == util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String("KtonBuildInEvent(address,address,uint256,bytes)")))) {
		currency = kton
	}

	return db.AddRingBurnRecord(Eth, util.AddHex(e.Result.TransactionHash), address, target, currency, amount,
		int(util.U256(e.Result.BlockNumber).Int64()), int(util.U256(e.Result.TimeStamp).Int64()))
}
