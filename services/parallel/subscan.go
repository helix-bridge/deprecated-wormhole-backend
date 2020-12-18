package parallel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/util"
	"strings"
)

type SubscanEventsRes struct {
	Data struct {
		Events []struct {
			BlockNum   int64  `json:"block_num"`
			Params     string `json:"params"`
			EventIndex string `json:"event_index"`
			EventId    string `json:"event_id"`
			ModuleId   string `json:"module_id"`
		} `json:"events"`
	} `json:"data"`
}

type SubscanExtrinsicRes struct {
	Data struct {
		BlockTimestamp int `json:"block_timestamp"`
		BlockNum       int `json:"block_num"`
		Events         []struct {
			BlockNum   int64  `json:"block_num"`
			Params     string `json:"params"`
			EventIndex string `json:"event_index"`
			EventId    string `json:"event_id"`
			ModuleId   string `json:"module_id"`
		} `json:"event"`
		BlockHash string `json:"block_hash"`
	} `json:"data"`
}

type SubscanEvent struct {
	BlockNum       int64        `json:"block_num"`
	Params         []EventParam `json:"event"`
	ExtrinsicIndex string       `json:"event_index"`
	EventId        string       `json:"event_id"`
	ModuleId       string       `json:"module_id"`
}

type EventParam struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type SubscanParams struct {
	Row       int    `json:"row"`
	Call      string `json:"call"`
	Module    string `json:"module"`
	FromBlock int64  `json:"from_block"`
}

func SubscanEvents(moduleId, eventId string, startBlock int64) (list []SubscanEvent) {
	var res SubscanEventsRes
	url := fmt.Sprintf("%s/api/scan/events", config.Link.SubscanHost)
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
		list = append(list, SubscanEvent{BlockNum: event.BlockNum, Params: params, ExtrinsicIndex: event.EventIndex, EventId: event.EventId, ModuleId: event.ModuleId})
	}
	return
}

type ExtrinsicDetail struct {
	BlockNum       int            `json:"block_num"`
	Event          []SubscanEvent `json:"event"`
	BlockHash      string         `json:"block_hash"`
	BlockTimestamp int            `json:"block_timestamp"`
}

func SubscanExtrinsic(extrinsicIndex string) *ExtrinsicDetail {
	var res SubscanExtrinsicRes
	url := fmt.Sprintf("%s/api/scan/extrinsic", config.Link.SubscanHost)
	raw, err := util.PostWithJson(url, strings.NewReader(fmt.Sprintf(`{"extrinsic_index":"%s"}`, extrinsicIndex)))

	if err != nil {
		return nil
	}
	util.UnmarshalAny(&res, raw)
	util.Debug(res)
	var detail ExtrinsicDetail
	detail.BlockNum = res.Data.BlockNum
	detail.BlockHash = res.Data.BlockHash
	detail.BlockTimestamp = res.Data.BlockTimestamp
	var list []SubscanEvent
	for _, event := range res.Data.Events {
		var params []EventParam
		util.UnmarshalAny(&params, event.Params)
		list = append(list, SubscanEvent{BlockNum: event.BlockNum, Params: params, ExtrinsicIndex: event.EventIndex, EventId: event.EventId, ModuleId: event.ModuleId})
	}
	detail.Event = list
	return &detail
}
