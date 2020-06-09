package parallel

import (
	"github.com/darwinia-network/link/lib/web3"
	"github.com/darwinia-network/link/util"
	"math/big"
)

var tronContract = "416e0d26adf5323f5b82d5714354dc3c6870adee7c"

type TronResponse struct {
	ConstantResult []string `json:"constant_result"`
}

func RingTronSupply() *big.Int {
	w := web3.New("tron")
	var e TronResponse
	_ = w.Call(&e, tronContract, "totalSupply()")
	return util.U256(e.ConstantResult[0])
}

type TronScan struct {
	Success bool             `json:"success"`
	Data    []TronScanResult `json:"data"`
}

type TronScanResult struct {
	TransactionId string            `json:"transaction_id"`
	EventName     string            `json:"event_name"`
	Result        map[string]string `json:"result"`
}

func TronScanLog(start int64, address string) (*TronScan, error) {
	w := web3.New("tron")
	var e TronScan
	if err := w.Event(&e, start, address, ""); err != nil || !e.Success {
		return nil, err
	}
	return &e, nil
}
