package db

import (
	"errors"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"strings"
	"time"
)

type TokenRegisterRecord struct {
	ExtrinsicIndex string          `json:"extrinsic_index" gorm:"primary_key;auto_increment:false"`
	CreatedAt      time.Time       `json:"-"`
	AccountId      string          `json:"account_id"`
	BlockNum       int             `json:"block_num"`
	BlockHash      string          `sql:"default: null;size:70" json:"block_hash" `
	Backing        string          `json:"backing" sql:"default: null;size:100"`
	Source         string          `json:"source" sql:"default: null;size:100"`
	Target         string          `json:"target" sql:"default: null;size:100"`
	BlockTimestamp int             `json:"block_timestamp"`
	MMRIndex       uint            `json:"mmr_index"`
	MMRRoot        string          `json:"mmr_root"`
	Signatures     string          `json:"signatures" sql:"type:text;"`
	BlockHeader    string          `json:"block_header" sql:"type:text;"`
	Tx             string          `json:"tx"`
}

type TokenBurnRecord struct {
	ExtrinsicIndex string          `json:"extrinsic_index" gorm:"primary_key;auto_increment:false"`
	CreatedAt      time.Time       `json:"-"`
	AccountId      string          `json:"account_id"`
	BlockNum       int             `json:"block_num"`
	BlockHash      string          `sql:"default: null;size:70" json:"block_hash" `
	Backing        string          `json:"backing" sql:"default: null;size:100"`
	Source         string          `json:"source" sql:"default: null;size:100"`
	Target         string          `json:"target" sql:"default: null;size:100"`
	Sender         string          `json:"sender" sql:"default: null;size:100"`
	Recipient      string          `json:"recipient" sql:"default: null;size:100"`
	Value          string          `json:"value" sql:"default: null;size:100"`
	BlockTimestamp int             `json:"block_timestamp"`
	MMRIndex       uint            `json:"mmr_index"`
	MMRRoot        string          `json:"mmr_root"`
	Signatures     string          `json:"signatures" sql:"type:text;"`
	BlockHeader    string          `json:"block_header" sql:"type:text;"`
	Tx             string          `json:"tx"`
}

func CreateTokenRegisterRecord(extrinsicIndex string, detail *parallel.ExtrinsicDetail) error {
	db := util.DB

	record := &TokenRegisterRecord{
		ExtrinsicIndex: extrinsicIndex,
		BlockHash:      detail.BlockHash,
		BlockNum:       detail.BlockNum,
		BlockTimestamp: detail.BlockTimestamp,
		BlockHeader:    util.ToString(parallel.SubscanBlockHeader(uint(detail.BlockNum))),
	}
	for _, event := range detail.Event {
		switch event.EventId {
		case "TokenRegistered":
			record.Backing = util.AddHex(util.ToString(event.Params[1].Value))
			record.Source  = util.AddHex(util.ToString(event.Params[2].Value))
			record.Target  = util.AddHex(util.ToString(event.Params[3].Value))
		case "ScheduleMMRRoot":
			record.MMRIndex = uint(util.StringToInt(util.ToString(event.Params[0].Value)))
		}
	}
	if record.MMRIndex == 0 {
		// latest MMRIndex
		var recent TokenRegisterRecord
		if query := db.Model(TokenRegisterRecord{}).Where("mmr_index > ", detail.BlockNum).Order("mmr_index asc").Limit(1).Find(&recent); !query.RecordNotFound() {
			record.MMRIndex = recent.MMRIndex
		} else {
			return errors.New("nil MMRIndex")
		}

	}

	query := db.Create(&record)
	return query.Error
}

func CreateTokenBurnRecord(extrinsicIndex string, detail *parallel.ExtrinsicDetail) error {
	db := util.DB

	record := &TokenBurnRecord{
		ExtrinsicIndex: extrinsicIndex,
		BlockHash:      detail.BlockHash,
		BlockNum:       detail.BlockNum,
		BlockTimestamp: detail.BlockTimestamp,
		BlockHeader:    util.ToString(parallel.SubscanBlockHeader(uint(detail.BlockNum))),
	}
	for _, event := range detail.Event {
		switch event.EventId {
		case "BurnToken":
			record.Backing = util.AddHex(util.ToString(event.Params[1].Value))
			record.Sender = util.AddHex(util.ToString(event.Params[2].Value))
			record.Recipient  = util.AddHex(util.ToString(event.Params[3].Value))
			record.Source  = util.AddHex(util.ToString(event.Params[4].Value))
			record.Target  = util.AddHex(util.ToString(event.Params[5].Value))
			record.Value  = util.AddHex(util.ToString(event.Params[6].Value))
		case "ScheduleMMRRoot":
			record.MMRIndex = uint(util.StringToInt(util.ToString(event.Params[0].Value)))
		}
	}
	if record.MMRIndex == 0 {
		// latest MMRIndex
		var recent TokenBurnRecord
		if query := db.Model(TokenBurnRecord{}).Where("mmr_index > ", detail.BlockNum).Order("mmr_index asc").Limit(1).Find(&recent); !query.RecordNotFound() {
			record.MMRIndex = recent.MMRIndex
		} else {
			return errors.New("nil MMRIndex")
		}

	}

	query := db.Create(&record)
	return query.Error
}

func MMRRootSignedForTokenRegistration(eventParams []parallel.EventParam) error {
	mmrIndex := util.IntFromInterface(eventParams[0].Value)
	mmrRoot := util.ToString(eventParams[1].Value)

	var signatures []MMRSignedSignature
	util.UnmarshalAny(&signatures, eventParams[2].Value)

	var signatureList []string
	for _, signature := range signatures {
		signatureList = append(signatureList, signature.Col2)
	}

	queryRegister := util.DB.Model(TokenRegisterRecord{}).Where("tx = ''").Where("block_num < ?", mmrIndex).
		Update(TokenRegisterRecord{MMRRoot: util.AddHex(mmrRoot), Signatures: strings.Join(signatureList, ","), MMRIndex: uint(mmrIndex)})
	queryBurned := util.DB.Model(TokenBurnRecord{}).Where("tx = ''").Where("block_num < ?", mmrIndex).
		Update(TokenBurnRecord{MMRRoot: util.AddHex(mmrRoot), Signatures: strings.Join(signatureList, ","), MMRIndex: uint(mmrIndex)})
	if queryRegister.Error != nil && queryBurned.Error != nil {
		return errors.New("operator db failed")
	}
	return nil
}

func SetTokenTokenRegistrationConfirm(blockNum uint64, tx string) error {
	queryRegister := util.DB.Model(TokenRegisterRecord{}).Where("block_num = ?", blockNum).Update(TokenRegisterRecord{Tx: tx})
	queryBurned := util.DB.Model(TokenBurnRecord{}).Where("block_num = ?", blockNum).Update(TokenBurnRecord{Tx: tx})
	if queryRegister.Error != nil && queryBurned.Error != nil {
		return errors.New("operator db failed")
	}
	return nil
}

func TokenRegisterRecordInfo(source string) *TokenRegisterRecord {
	var r TokenRegisterRecord
	query := util.DB.Model(TokenRegisterRecord{}).Where("source = ?", source).Find(&r)
	if query.Error != nil || query.RecordNotFound() {
		return nil
	}
	return &r
}

func (r *TokenBurnRecord) setMMRRoot(mmrRoot string) error {
	query := util.DB.Model(r).Update(TokenBurnRecord{MMRRoot: mmrRoot})
	return query.Error
}

func TokenBurnRecords(sender string, page, row int) ([]TokenBurnRecord, int) {
	var list []TokenBurnRecord
	var count int
	util.DB.Model(TokenBurnRecord{}).Where("sender = ?", sender).Count(&count)
	util.DB.Where("sender = ?", sender).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)

	for index, burned := range list {
		if burned.MMRRoot != "" {
			continue
		}
		if MMRRoot := queryMMRRoot(burned.MMRIndex + 1); MMRRoot != "" {
			_ = burned.setMMRRoot(MMRRoot)
			list[index].MMRRoot = MMRRoot
		}
	}
	return list, count
}

