package parallel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/util"
	"github.com/darwinia-network/link/util/log"
	"io"
	"io/ioutil"
	"net/http"
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

type SubscanBlockHeaderRes struct {
	Data *BlockHeader `json:"data"`
}

type BlockHeader struct {
	ParentHash     string   `json:"parent_hash"`
	BlockNumber    int      `json:"block_number"`
	StateRoot      string   `json:"state_root"`
	ExtrinsicsRoot string   `json:"extrinsics_root"`
	Digest         []string `json:"digest"`
}

type SubscanLogsRes struct {
	Data []SubscanLog `json:"data"`
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
	TypeName string   `json:"type_name"`
	Value interface{} `json:"value"`
}

type SubscanParams struct {
	Row       int    `json:"row"`
	Page      int64    `json:"page"`
	Call      string `json:"call"`
	Module    string `json:"module"`
	FromBlock int64  `json:"from_block"`
	Finalized bool   `json:"finalized"`
}

type SubscanLog struct {
	LogType string `json:"log_type"`
	Data    string `json:"data"`
}

func SubscanEvents(moduleId, eventId string, startBlock int64, page int64) (list []SubscanEvent) {
	var res SubscanEventsRes
	url := fmt.Sprintf("%s/api/scan/events", config.Link.SubscanHost)
	p := SubscanParams{
		Row:       100,
		Page:      page,
		Call:      eventId,
		Module:    moduleId,
		FromBlock: startBlock,
		Finalized: true,
	}
	bp, _ := json.Marshal(p)
	raw, err := PostWithApiKey(url, bytes.NewReader(bp))
	if err != nil {
		log.Error("post subscan event failed", "err", err)
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
	raw, err := PostWithApiKey(url, strings.NewReader(fmt.Sprintf(`{"extrinsic_index":"%s"}`, extrinsicIndex)))

	if err != nil {
		return nil
	}
	util.UnmarshalAny(&res, raw)
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

func SubscanLogs(blockNum uint) []SubscanLog {
	var res SubscanLogsRes
	url := fmt.Sprintf("%s/api/scan/logs", config.Link.SubscanHost)
	raw, err := PostWithApiKey(url, strings.NewReader(fmt.Sprintf(`{"block_num":%d}`, blockNum)))
	if err != nil {
		return nil
	}
	util.UnmarshalAny(&res, raw)
	return res.Data
}

func SubscanBlockHeader(blockNum uint) *BlockHeader {
	var res SubscanBlockHeaderRes
	url := fmt.Sprintf("%s/api/scan/header", config.Link.SubscanHost)
	raw, err := PostWithApiKey(url, strings.NewReader(fmt.Sprintf(`{"block_num":%d}`, blockNum)))
	if err != nil {
		return nil
	}
	util.UnmarshalAny(&res, raw)
	return res.Data
}

func PostWithApiKey(url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", util.GetEnv("SUBSCAN_API_KEY", "16325b647dbe652fd9fa73ceaaa8aa83"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, fmt.Errorf("empty response")
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
