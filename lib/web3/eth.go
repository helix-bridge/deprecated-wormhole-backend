package web3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/crypto"
	"net/url"
)

type eth struct{}

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

type Transaction struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash   string `json:"blockHash"`
		BlockNumber string `json:"blockNumber"`
		From        string `json:"from"`
		Gas         string `json:"gas"`
		GasPrice    string `json:"gasPrice"`
		Hash        string `json:"hash"`
	} `json:"result"`
}

func (e *eth) url() string {
	return util.GetEnv("INFURA", "https://ropsten.infura.io/v3/67fb2d92380a4ff8a2b1ebef24a81a8f")
}

func (e *eth) scan() string {
	return fmt.Sprintf("%s/api?module=logs&action=getLogs&", util.GetEnv("ETHERSCAN", "https://api-ropsten.etherscan.io"))
}

func (e *eth) Call(v interface{}, contract, method string, params ...string) error {
	sha3Function := util.AddHex(util.BytesToHex(crypto.SoliditySHA3(crypto.String(method))))
	body := make([]interface{}, 2)
	var data string
	for _, param := range params {
		data = data + util.Padding(param)
	}
	body[0] = TransactionParameters{
		To:   contract,
		Data: sha3Function[0:10] + data,
	}
	body[1] = "latest"
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
	return json.Unmarshal(response, v)
}

// only support two topic query
func (e *eth) Event(v interface{}, start int64, address string, topic ...string) error {
	if len(topic) < 1 {
		return fmt.Errorf("need at least one topic")
	}
	etherscan := e.scan()
	q := url.Values{}
	q.Add("fromBlock", util.Int64ToString(start))
	q.Add("toBlock", "latest")
	q.Add("address", address)
	q.Add("topic0", topic[0])
	// if len(topic) == 2 {
	// 	q.Add("topic0_1_opr", "or")
	// 	q.Add("topic1", topic[1])
	// }
	q.Add("apikey", util.GetEnv("ETHSCAN_KEY", "G78R6SGMHGXSMXZCBDW8WE716YQFQGJ68F"))
	fmt.Println(fmt.Sprintf("%s%s", etherscan, q.Encode()))
	response, err := util.HttpGet(fmt.Sprintf("%s%s", etherscan, q.Encode()))
	if err != nil {
		return err
	}
	return json.Unmarshal(response, v)
}

func (e *eth) GetTransactionByBlockHashAndIndex(blockHash string, index int) string {
	r := ReqBody{
		JSONRPC: "2.0",
		Method:  "eth_getTransactionByBlockHashAndIndex",
		ID:      1,
		Params:  []interface{}{blockHash, fmt.Sprintf("0x%x", index)},
	}
	j, _ := json.Marshal(r)
	b, err := util.PostWithJson(e.url(), bytes.NewReader(j))
	if b == nil || err != nil {
		return ""
	}
	var transaction Transaction
	err = json.Unmarshal(b, &transaction)
	if err != nil {
		return ""
	}
	return transaction.Result.Hash
}
