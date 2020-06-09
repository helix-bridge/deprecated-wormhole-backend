package web3

import (
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/util"
	"net/url"
	"strings"
	"time"
)

type tron struct {
}

func (e *tron) url() string {
	return "https://api.trongrid.io"
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

// address=TSbJFbH8sSayRFMavwohY2P6QfKwQEWcaz&start=0&limit=200&start_timestamp=1548000000000&end_timestamp=1548056638507
func (e *tron) Event(v interface{}, start int64, address, topic string) error {
	etherscan := "https://api.shasta.tronscan.org/api/contract/events?"
	q := url.Values{}
	q.Add("address", address)
	q.Add("start", "0")
	q.Add("limit", "200")
	q.Add("start_timestamp", util.Int64ToString(start*1000))
	q.Add("end_timestamp", util.Int64ToString(time.Now().Unix()*1000))
	fmt.Println(fmt.Sprintf("%s%s", etherscan, q.Encode()))
	response, err := util.HttpGet(fmt.Sprintf("%s?%s", etherscan, q.Encode()))
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return json.Unmarshal(response, v)
}
