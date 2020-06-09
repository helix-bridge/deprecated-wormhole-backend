package observer

import (
	"github.com/darwinia-network/link/services/parallel"
	"time"
)

type TronTransaction struct {
	Last    int64       `json:"last"`
	Address string      `json:"address"`
	Method  string      `json:"method"`
	Result  interface{} `json:"result"`
}

func (e *TronTransaction) Do(o Observable) error {
	return nil
}

func (e *TronTransaction) Listen(o Observable) error {
	go func() {
		for {
			if eventLog, _ := parallel.TronScanLog(e.Last, e.Address, e.Address); eventLog != nil {
				for _, result := range eventLog.Result {
					e.Result = result
					_ = o.notify(e)
				}
				e.Last = time.Now().Unix()
			}
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}
