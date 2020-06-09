package observer

import "runtime"

type Observable interface {
	Attach(observer ...IObserver) Observable
	notify(observer IObserver) error
	Run() error
}

type IObserver interface {
	Do(o Observable) error
	Listen(o Observable) error
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

func Run() {
	subject := &ObservableConcrete{}
	subject.Attach(&EthTransaction{Address: "0xb88141E0B8702f9bcFd8063c2Ac8852771525c4e", Method: "RingBuildInEvent(address,address,uint256,bytes)"})
	_ = subject.Run()
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
