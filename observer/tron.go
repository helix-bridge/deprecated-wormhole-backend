package observer

import (
	"fmt"
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type TronTransaction struct {
	Last    int64                    `json:"last"`
	Address string                   `json:"address"`
	Method  []string                 `json:"method"`
	Result  *parallel.TronScanResult `json:"result"`
}

func (e *TronTransaction) Do(o Observable) error {
	fmt.Println("find TronTransaction", e.Result)
	if e.Result == nil || !util.StringInSlice(e.Result.EventName, TronAvailableEvent) {
		return fmt.Errorf("empty transaction")
	}
	return e.RingBurnRecord()
}

func (e *TronTransaction) Listen(o Observable) error {
	key := runFuncName()
	if e.Last == 0 {
		if b := util.GetCache(key); b != nil {
			e.Last = util.StringToInt64(string(b))
		} else {
			e.Last = 1591683963
		}
	}
	go func() {
		for {
			if eventLog, _ := parallel.TronScanLog(e.Last, e.Address); eventLog != nil {
				for _, result := range eventLog.Data {
					if util.StringInSlice(result.EventName, e.Method) {
						e.Result = &result
						_ = o.notify(e)
					}
				}
				e.Last = time.Now().Unix()
				_ = util.SetCache(key, e.Last, 86400*7)
			}
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}

func (e *TronTransaction) RingBurnRecord() error {
	address := util.TrimHex(e.Result.Result["owner"])
	amount := decimal.RequireFromString(e.Result.Result["amount"])
	target := e.Result.Result["data"]

	currency := ring
	if e.Result.EventName == "KtonBuildInEvent" {
		currency = kton
	}

	return db.AddRingBurnRecord(Tron, util.AddHex(e.Result.TransactionId), util.AddTronPerfix(address), target, currency, amount, e.Result.BlockNumber)
}
