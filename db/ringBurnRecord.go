package db

import (
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type RingBurnRecord struct {
	ID             uint            `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time       `json:"created_at"`
	Chain          string          `json:"chain"`
	Tx             string          `json:"tx" sql:"size:100"`
	Address        string          `json:"address" sql:"size:100"`
	Target         string          `json:"target" sql:"size:100"`
	Currency       string          `json:"currency"`
	BlockNum       int             `json:"block_num"`
	Amount         decimal.Decimal `json:"amount" sql:"type:decimal(40,0);" `
	BlockTimestamp int             `json:"block_timestamp"`
}

func AddRingBurnRecord(chain, tx, address, target, currency string, amount decimal.Decimal, blockNum, blockTimestamp int) error {
	db := util.DB
	query := db.Create(&RingBurnRecord{
		Chain: chain, Tx: tx, Address: address, Target: target, Amount: amount, Currency: currency, BlockNum: blockNum,
		BlockTimestamp: blockTimestamp,
	})
	return query.Error
}

func RingBurnList(address string, page, row int) ([]RingBurnRecord, int) {
	db := util.DB
	var list []RingBurnRecord
	var count int
	db.Model(RingBurnRecord{}).Where("address = ?", address).Count(&count)
	if row > 0 {
		db.Where("address = ?", address).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)
	} else {
		db.Model(list).Where("address = ?", address).Order("block_num desc").Find(&list)
	}
	return list, count
}
