package observer

import (
	"fmt"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"time"
)

type TronTransaction struct {
	Last    int64       `json:"last"`
	Address string      `json:"address"`
	Method  string      `json:"method"`
	Result  interface{} `json:"result"`
}

func (e *TronTransaction) Do(o Observable) error {
	fmt.Println("find TronTransaction", e.Result)
	return nil
}

func (e *TronTransaction) Listen(o Observable) error {
	key := runFuncName()
	if e.Last == 0 {
		if b := util.GetCache(key); b != nil {
			e.Last = util.StringToInt64(string(b))
		} else {
			e.Last = 1591683963000
		}
	}
	go func() {
		for {
			if eventLog, _ := parallel.TronScanLog(e.Last, e.Address); eventLog != nil {
				for _, result := range eventLog.Data {
					if result.EventName == e.Method {
						e.Result = result
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
