package web3

import (
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/util"
	"strings"
)

type tron struct {
}

func (e *tron) url() string {
	return "https://api.trongrid.io"
}

func (e *tron) Call(v interface{}, contract, method string, params ...string) error {
	url := e.url() + "/wallet/triggersmartcontract"
	body := fmt.Sprintf("{\"contract_address\":\"%s\",\"function_selector\":\"%s\",\"fee_limit\":%d,\"call_value\":%d,\"owner_address\":\"%s\",\"parameter\":\"",
		contract, method, 0, 0, contract)
	for _, p := range params {
		body += util.Padding(p)
	}

	body += "\"}"

	response, err := util.PostWithJson(url, strings.NewReader(body))
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return json.Unmarshal(response, v)
}
