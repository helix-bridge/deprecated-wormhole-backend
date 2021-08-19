package web3

type Web3 interface {
	Call(v interface{}, contract, method string, params ...string) error
	url() string
	Event(v interface{}, start int64, to int64, address string, topic ...string) error
	GetTransactionByBlockHashAndIndex(blockHash string, index int) string
}

type web3 struct{}

func (w3 *web3) GetTransactionByBlockHashAndIndex(blockHash string, index int) string {
	panic("implement me")
}

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

func (w3 *web3) Event(v interface{}, start, to int64, address string, topic ...string) error {
	return nil
}
