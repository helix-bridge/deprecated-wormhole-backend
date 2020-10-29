package parallel

import (
	"bytes"
	"encoding/json"
	"github.com/darwinia-network/link/util"
)

type SubscanEventsRes struct {
	Data struct {
		Events []struct {
			BlockNum   int64  `json:"block_num"`
			Params     string `json:"params"`
			EventIndex string `json:"event_index"`
		} `json:"events"`
	} `json:"data"`
}

type SubscanEvent struct {
	BlockNum   int64        `json:"block_num"`
	Params     []EventParam `json:"params"`
	EventIndex string       `json:"event_index"`
}

type EventParam struct {
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"valueRaw"`
}

type SubscanParams struct {
	Row       int    `json:"row"`
	Call      string `json:"call"`
	Module    string `json:"module"`
	FromBlock int64  `json:"from_block"`
}

func SubscanEvents(moduleId, eventId string, startBlock int64) (list []SubscanEvent) {
	var res SubscanEventsRes
	url := "https://crab.subscan.io/api/scan/events"
	p := SubscanParams{
		Row:       100,
		Call:      eventId,
		Module:    moduleId,
		FromBlock: startBlock,
	}
	bp, _ := json.Marshal(p)
	raw, err := util.PostWithJson(url, bytes.NewReader(bp))
	if err != nil {
		return nil
	}
	util.UnmarshalAny(&res, raw)
	for _, event := range res.Data.Events {
		var params []EventParam
		util.UnmarshalAny(&params, event.Params)
		list = append(list, SubscanEvent{
			BlockNum:   event.BlockNum,
			Params:     params,
			EventIndex: event.EventIndex,
		})
	}
	return
}
