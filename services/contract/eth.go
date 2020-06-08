package contract

import (
	"github.com/darwinia-network/link/lib/web3"
	"github.com/darwinia-network/link/util"
	"math/big"
)

var ethContract = "0x9469d013805bffb7d3debe5e7839237e535ec483"

type EthResponse struct {
	Result string `json:"result,omitempty"`
}

func RingEthSupply() *big.Int {
	w := web3.New("eth")
	var e EthResponse
	_ = w.Call(&e, ethContract, "totalSupply()")
	return util.U256(e.Result)
}
