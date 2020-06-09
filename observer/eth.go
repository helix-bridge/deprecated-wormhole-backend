package observer

import (
	"fmt"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"time"
)

type EthTransaction struct {
	Last    int64                     `json:"last"`
	Address string                    `json:"address"`
	Method  string                    `json:"method"`
	Result  *parallel.EtherscanResult `json:"result"`
}

func (e *EthTransaction) Do(o Observable) error {
	fmt.Println("find EthTransaction", e.Result)
	return nil
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
			if eventLog, _ := parallel.EtherscanLog(e.Last+1, e.Address, e.Method); eventLog != nil {
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
