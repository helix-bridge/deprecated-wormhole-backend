package observer

import (
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/util"
	"runtime"
	"time"
)

type Observable interface {
	Attach(observer ...IObserver) Observable
	notify(observer IObserver) error
	Run() error
}

type IObserver interface {
	Do(o Observable) error
	Listen(o Observable) error
	Pause()
	Resume()
	ErrorBreak(error)
	LoadData(o Observable, isRely bool)
}

type ObservableConcrete struct {
	observerList []IObserver
}

func (o *ObservableConcrete) Attach(observer ...IObserver) Observable {
	o.observerList = append(o.observerList, observer...)
	return o
}

func (o *ObservableConcrete) notify(observer IObserver) (err error) {
	for _, item := range o.observerList {
		if item == observer {
			if err = observer.Do(o); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *ObservableConcrete) Run() (err error) {
	for _, item := range o.observerList {
		if err = item.Listen(o); err != nil {
			return err
		}
	}
	return nil
}

func (o *ObservableConcrete) Pause() {
	for _, item := range o.observerList {
		item.Pause()
	}
}

func (o *ObservableConcrete) Resume() {
	for _, item := range o.observerList {
		item.Resume()
	}
}

func (o *ObservableConcrete) Monitor() {
	monitorInterval := time.Second * time.Duration(10)
	monitorTimer := time.NewTimer(monitorInterval)
	go func() {
		for {
			select {
			case <-monitorTimer.C:
				restart := util.HgetCache("restart", "restart")
				if string(restart) == "true" {
					o.Pause()
					for _, item := range o.observerList {
						item.LoadData(o, false)
					}
					for _, item := range o.observerList {
						item.LoadData(o, true)
					}
					util.HsetCache("restart", "restart", []byte("false"))
					o.Resume()
				}
				monitorTimer.Reset(monitorInterval)
			}
		}
	}()
}

func Run() {
	subject := &ObservableConcrete{}

	subject.Attach(
		&EthTransaction{Address: config.Link.TokenRedeem, Method: []string{"BurnAndRedeem(address,address,uint256,bytes)"}},
		&EthTransaction{Address: config.Link.DepositRedeem, Method: []string{"BurnAndRedeem(uint256,address,uint48,uint48,uint64,uint128,bytes)"}},
		&EthTransaction{Address: config.Link.TokenIssuing, Method: []string{"VerifyProof(uint32)"}},
		&EthTransaction{Address: config.Link.EthBridgerRelay, Method: []string{"SetRootEvent(address,bytes32,uint256)"}},
		&EthTransaction{Address: config.Link.EthereumBacking, Method: []string{"VerifyProof(uint32)"}},
		&EthTransaction{Address: config.Link.EthereumBacking, Method: []string{"BackingLock(address,address,address,uint256,address,uint256)"}},
		&SubscanEvent{ModuleId: "ethereumrelay", EventId: "PendingRelayHeaderParcelConfirmed"},
		&SubscanEvent{ModuleId: "ethereumbacking"},
		&SubscanEvent{ModuleId: "ethereumrelayauthorities", EventId: "MMRRootSigned"},
		&SubscanEvent{ModuleId: "ethereumissuing"},
	)
	subject.Monitor()
	_ = subject.Run()
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
