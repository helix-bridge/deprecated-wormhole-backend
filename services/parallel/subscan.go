package parallel

import (
	"fmt"
	"github.com/darwinia-network/link/util"
	"strings"
)

type SubscanEventsRes struct {
	Data struct {
		Events []struct {
			BlockNum int    `json:"block_num"`
			Params   string `json:"params"`
		} `json:"events"`
	} `json:"data"`
}

type SubscanEvent struct {
	BlockNum int          `json:"block_num"`
	Params   []EventParam `json:"params"`
}

type EventParam struct {
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"valueRaw"`
}

func SubscanEvents(eventId string) (list []SubscanEvent) {
	var res SubscanEventsRes
	url := "https://crab.subscan.io/api/scan/events"
	raw, err := util.PostWithJson(url, strings.NewReader(fmt.Sprintf(`{"row": 1, "page": 0, "call": "%s"}`, eventId)))
	if err != nil {
		return nil
	}
	util.UnmarshalAny(&res, raw)
	for _, event := range res.Data.Events {
		var params []EventParam
		util.UnmarshalAny(&params, event.Params)
		list = append(list, SubscanEvent{
			BlockNum: event.BlockNum,
			Params:   params,
		})
	}
	return
}
