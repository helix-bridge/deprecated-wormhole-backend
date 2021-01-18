package db

import (
	"errors"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type DarwiniaBackingLock struct {
	ExtrinsicIndex string          `json:"extrinsic_index" gorm:"primary_key;auto_increment:false"`
	CreatedAt      time.Time       `json:"-"`
	AccountId      string          `json:"account_id"`
	BlockNum       int             `json:"block_num"`
	BlockHash      string          `sql:"default: null;size:70" json:"block_hash" `
	RingValue      decimal.Decimal `json:"ring_value" sql:"type:decimal(30,0);"`
	KtonValue      decimal.Decimal `json:"kton_value" sql:"type:decimal(30,0);"`
	Target         string          `json:"target" sql:"default: null;size:100"`
	BlockTimestamp int             `json:"block_timestamp"`
	MMRIndex       uint            `json:"mmr_index"`
	MMRRoot        string          `json:"mmr_root"`
	Signatures     string          `json:"signatures" sql:"type:text;"`
	BlockHeader    string          `json:"block_header" sql:"type:text;"`
	Tx             string          `json:"tx"`
}

func CreateDarwiniaBacking(extrinsicIndex string, detail *parallel.ExtrinsicDetail) error {
	db := util.DB

	record := &DarwiniaBackingLock{
		ExtrinsicIndex: extrinsicIndex,
		BlockHash:      detail.BlockHash,
		BlockNum:       detail.BlockNum,
		BlockTimestamp: detail.BlockTimestamp,
		BlockHeader:    util.ToString(parallel.SubscanBlockHeader(uint(detail.BlockNum))),
	}
	for _, event := range detail.Event {
		switch event.EventId {
		case "LockRing":
			record.AccountId = util.AddHex(util.ToString(event.Params[0].Value))
			record.Target = util.ToString(event.Params[1].Value)
			record.RingValue = util.DecimalFromInterface(event.Params[3].Value)
		case "LockKton":
			record.AccountId = util.AddHex(util.ToString(event.Params[0].Value))
			record.Target = util.ToString(event.Params[1].Value)
			record.KtonValue = util.DecimalFromInterface(event.Params[3].Value)
		case "ScheduleMMRRoot":
			record.MMRIndex = uint(util.StringToInt(util.ToString(event.Params[0].Value)))
		}
	}
	if record.MMRIndex == 0 {
		// latest MMRIndex
		var recent DarwiniaBackingLock
		if query := db.Model(DarwiniaBackingLock{}).Where("mmr_index > ", detail.BlockNum).Order("mmr_index asc").Limit(1).Find(&recent); !query.RecordNotFound() {
			record.MMRIndex = recent.MMRIndex
		} else {
			return errors.New("nil MMRIndex")
		}

	}

	query := db.Create(&record)
	return query.Error
}

type MerkleMountainRangeRootLog struct {
	ParentMmrRoot string `json:"parent_mmr_root"`
}

func DarwiniaBackingLocks(accountId string, page, row int) ([]DarwiniaBackingLock, int) {
	var list []DarwiniaBackingLock
	var count int
	util.DB.Model(DarwiniaBackingLock{}).Where("account_id = ?", accountId).Count(&count)
	util.DB.Where("account_id = ?", accountId).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)

	for index, lock := range list {
		if lock.MMRRoot != "" {
			continue
		}
		if MMRRoot := queryMMRRoot(lock.MMRIndex + 1); MMRRoot != "" {
			_ = lock.setMMRRoot(MMRRoot)
			list[index].MMRRoot = MMRRoot
		}
	}
	return list, count
}

func queryMMRRoot(blockNum uint) string {
	logs := parallel.SubscanLogs(blockNum)

	for _, logData := range logs {
		if strings.EqualFold(logData.LogType, "other") {

			var merkleMountainRangeRootLog *MerkleMountainRangeRootLog
			if util.UnmarshalAny(&merkleMountainRangeRootLog, logData.Data); merkleMountainRangeRootLog != nil {
				return merkleMountainRangeRootLog.ParentMmrRoot
			}
		}
	}
	return ""
}

func (l *DarwiniaBackingLock) setMMRRoot(mmrRoot string) error {
	query := util.DB.Model(l).Update(DarwiniaBackingLock{MMRRoot: mmrRoot})
	return query.Error
}

type MMRSignedSignature struct {
	Col2 string `json:"col2"`
}

func MMRRootSigned(eventParams []parallel.EventParam) error {
	mmrIndex := util.IntFromInterface(eventParams[0].Value)
	mmrRoot := util.ToString(eventParams[1].Value)

	var signatures []MMRSignedSignature
	util.UnmarshalAny(&signatures, eventParams[2].Value)

	var signatureList []string
	for _, signature := range signatures {
		signatureList = append(signatureList, signature.Col2)
	}

	query := util.DB.Model(DarwiniaBackingLock{}).Where("tx = ''").Where("block_num < ?", mmrIndex).
		Update(DarwiniaBackingLock{MMRRoot: util.AddHex(mmrRoot), Signatures: strings.Join(signatureList, ","), MMRIndex: uint(mmrIndex)})
	return query.Error
}

func SetBackingLockConfirm(blockNum uint64, tx string) error {
	query := util.DB.Model(DarwiniaBackingLock{}).Where("block_num = ?", blockNum).Update(DarwiniaBackingLock{Tx: tx})
	return query.Error
}

func BackingLock(extrinsicIndex string) *DarwiniaBackingLock {
	var d DarwiniaBackingLock
	query := util.DB.Model(DarwiniaBackingLock{}).Where("extrinsic_index = ?", extrinsicIndex).Find(&d)
	if query.Error != nil || query.RecordNotFound() {
		return nil
	}
	return &d
}

func SetMMRIndexBestBlockNum(blockNum uint64) {
	best, _ := GetMMRIndexBestBlockNum()
	if blockNum > best {
		_ = util.SetCache("MMRIndexBestBlockNum", blockNum, 86400*180)
		_ = util.SetCache("MMRIndexBestMMRRoot", queryMMRRoot(uint(blockNum+1)), 86400*180)
	}
}

func GetMMRIndexBestBlockNum() (uint64, string) {
	best := util.GetCacheUint64("MMRIndexBestBlockNum")
	if MMRRoot := string(util.GetCache("MMRIndexBestMMRRoot")); MMRRoot == "" {
		MMRRoot = queryMMRRoot(uint(best + 1))
		return best, MMRRoot
	} else {
		return best, MMRRoot
	}
}
