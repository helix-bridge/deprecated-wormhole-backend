package web3

import (
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/util"
	"net/url"
	"strings"
)

type tron struct {
}

func (e *tron) GetTransactionByBlockHashAndIndex(blockHash string, index int) string {
	return ""
}

func (e *tron) url() string {
	return "https://api.trongrid.io"
}

func (e *tron) scan() string {
	return util.GetEnv("TRONSCAN", "https://api.shasta.trongrid.io")
}

func (e *tron) Call(v interface{}, contract, method string, params ...string) error {
	body := fmt.Sprintf("{\"contract_address\":\"%s\",\"function_selector\":\"%s\",\"fee_limit\":%d,\"call_value\":%d,\"owner_address\":\"%s\",\"parameter\":\"",
		contract, method, 0, 0, contract)
	for _, p := range params {
		body += util.Padding(p)
	}

	body += "\"}"

	response, err := util.PostWithJson(e.url()+"/wallet/triggersmartcontract", strings.NewReader(body))
	if err != nil {
		return err
	}
	return json.Unmarshal(response, v)
}

func (e *tron) Event(v interface{}, start, _to int64, address string, topic ...string) error {
	trongrid := fmt.Sprintf("%s/v1/contracts/%s/events?", e.scan(), address)
	q := url.Values{}
	q.Add("only_confirmed", "true")
	q.Add("order_by", "block_timestamp,desc")
	q.Add("min_block_timestamp", util.Int64ToString(start*1000))
	fmt.Println(fmt.Sprintf("%s%s", trongrid, q.Encode()))
	response, err := util.HttpGet(fmt.Sprintf("%s?%s", trongrid, q.Encode()))
	if err != nil {
		return err
	}

	return json.Unmarshal(response, v)
}
