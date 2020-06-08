package web3

type Web3 interface {
	Call(v interface{}, contract, method string, params ...string) error
	url() string
}

type web3 struct{}

func New(chain string) Web3 {
	var w3 Web3

	switch chain {
	case "eth":
		w3 = &eth{}
	case "tron":
		w3 = &tron{}
	default:
		w3 = &web3{}
	}
	return w3
}

func (w3 *web3) Call(v interface{}, contract, method string, params ...string) error {
	return nil
}

func (w3 *web3) url() string {
	return ""
}
