package observer

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"time"
)

type SubscanEvent struct {
	EventId string
	Result  *parallel.SubscanEvent
}

func (s *SubscanEvent) Do(o Observable) error {
	return s.RelayBlock()
}

func (s *SubscanEvent) Listen(o Observable) error {
	go func() {
		for {
			if eventLog := parallel.SubscanEvents(s.EventId); eventLog != nil {
				for _, result := range eventLog {
					s.Result = &result
					_ = o.notify(s)
				}
			}
			time.Sleep(15 * time.Second)
		}
	}()
	return nil
}

func (s *SubscanEvent) RelayBlock() error {
	switch s.EventId {
	case "PendingRelayHeaderParcelApproved":
		db.SetRelayBestBlockNum(util.UInt64FromInterface(s.Result.Params[0].Value))
	}
	return nil
}
