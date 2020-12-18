package db

import (
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type DarwiniaBackingLock struct {
	ExtrinsicIndex string          `json:"extrinsic_Index" gorm:"primary_key;auto_increment:false"`
	CreatedAt      time.Time       `json:"-"`
	AccountId      string          `json:"account_id"`
	BlockNum       int             `json:"block_num"`
	BlockHash      string          `sql:"default: null;size:70" json:"block_hash" `
	RingValue      decimal.Decimal `json:"ring_value" sql:"type:decimal(30,0);"`
	KtonValue      decimal.Decimal `json:"kton_value" sql:"type:decimal(30,0);"`
	Target         string          `json:"target" sql:"default: null;size:100"`
	BlockTimestamp int             `json:"block_timestamp"`
}

func CreateDarwiniaBacking(extrinsicIndex string, detail *parallel.ExtrinsicDetail) error {
	db := util.DB

	record := &DarwiniaBackingLock{
		ExtrinsicIndex: extrinsicIndex,
		BlockHash:      detail.BlockHash,
		BlockNum:       detail.BlockNum,
		BlockTimestamp: detail.BlockTimestamp,
	}
	for _, event := range detail.Event {
		switch event.EventId {
		case "LockRing":
			record.AccountId = util.ToString(event.Params[0].Value)
			record.Target = util.ToString(event.Params[1].Value)
			record.RingValue = util.DecimalFromInterface(event.Params[3].Value)
		case "LockKton":
			record.AccountId = util.ToString(event.Params[0].Value)
			record.Target = util.ToString(event.Params[1].Value)
			record.KtonValue = util.DecimalFromInterface(event.Params[3].Value)
		}
	}

	query := db.Create(&record)
	return query.Error
}

func DarwiniaBackingLocks(accountId string, page, row int) ([]DarwiniaBackingLock, int) {
	var list []DarwiniaBackingLock
	var count int
	util.DB.Model(DarwiniaBackingLock{}).Where("account_id = ?", accountId).Count(&count)
	util.DB.Where("account_id = ?", accountId).Offset(page * row).Limit(row).Find(&list)
	return list, count
}
