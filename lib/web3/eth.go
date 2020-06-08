package web3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/crypto"
	"strings"
)

type eth struct {
}

type ReqBody struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	ID      int           `json:"id"`
	Params  []interface{} `json:"params"`
}

type TransactionParameters struct {
	To   string `json:"to"`
	Data string `json:"data,omitempty"`
}

func (e *eth) url() string {
	return "https://mainnet.infura.io/v3/1bb85682d6494e219803bab49a4813dc"
}

func (e *eth) Call(v interface{}, contract, method string, params ...string) error {
	sha3Function := util.BytesToHex(crypto.SoliditySHA3(crypto.String(method)))
	body := make([]interface{}, 2)
	body[0] = TransactionParameters{
		To:   contract,
		Data: util.AddHex(sha3Function[0:10] + strings.Join(params, "")),
	}
	body[1] = "latest"
	fmt.Println(body[0])
	r := ReqBody{
		JSONRPC: "2.0",
		Method:  "eth_call",
		ID:      1,
		Params:  body,
	}
	j, _ := json.Marshal(r)
	response, err := util.PostWithJson(e.url(), bytes.NewReader(j))
	if err != nil {
		return err
	}
	fmt.Println(string(response), string(j))
	return json.Unmarshal(response, v)
}
