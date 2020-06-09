package web3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/crypto"
	"net/url"
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
	return fmt.Sprintf("https://mainnet.infura.io/v3/%s", util.GetEnv("INFURA", "1bb85682d6494e219803bab49a4813dc"))
}

func (e *eth) Call(v interface{}, contract, method string, params ...string) error {
	sha3Function := util.BytesToHex(crypto.SoliditySHA3(crypto.String(method)))
	body := make([]interface{}, 2)
	body[0] = TransactionParameters{
		To:   contract,
		Data: util.AddHex(sha3Function[0:10] + strings.Join(params, "")),
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

func (e *eth) Event(v interface{}, start int64, address, topic string) error {
	etherscan := "https://api-ropsten.etherscan.io/api?module=logs&action=getLogs&"
	q := url.Values{}
	q.Add("fromBlock", util.Int64ToString(start))
	q.Add("toBlock", "latest")
	q.Add("address", address)
	q.Add("topic0", topic)
	q.Add("apikey", util.GetEnv("ETHSCAN_KEY", "G78R6SGMHGXSMXZCBDW8WE716YQFQGJ68F"))
	fmt.Println(fmt.Sprintf("%s%s", etherscan, q.Encode()))
	response, err := util.HttpGet(fmt.Sprintf("%s?%s", etherscan, q.Encode()))
	if err != nil {
		return err
	}
	return json.Unmarshal(response, v)
}
