package parallel

import (
	"encoding/hex"
	"github.com/darwinia-network/link/lib/web3"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/crypto"
	"math/big"
)

type Eth struct {
}

var ethContract = "0x9469d013805bffb7d3debe5e7839237e535ec483"

type EthResponse struct {
	Result string `json:"result,omitempty"`
}

type Etherscan struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Result  []EtherscanResult `json:"result"`
}

type EtherscanResult struct {
	Topics          []string `json:"topics"`
	Data            string   `json:"data"`
	TransactionHash string   `json:"transactionHash"`
	BlockNumber     string   `json:"blockNumber"`
	TimeStamp       string   `json:"timeStamp"`
}

func RingEthSupply() *big.Int {
	w := web3.New("eth")
	var e EthResponse
	_ = w.Call(&e, ethContract, "totalSupply()")
	return util.U256(e.Result)
}

func EtherscanLog(start int64, address string, methods ...string) (*Etherscan, error) {
	w := web3.New("eth")
	var e Etherscan
	var topics []string
	for _, method := range methods {
		topics = append(topics, util.AddHex(hex.EncodeToString(crypto.SoliditySHA3(crypto.String(method)))))
	}
	if err := w.Event(&e, start, address, topics...); err != nil || e.Message != "OK" {
		return nil, err
	}
	return &e, nil
}
