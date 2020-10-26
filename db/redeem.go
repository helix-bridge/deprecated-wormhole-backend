package db

import (
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type RedeemRecord struct {
	ID             uint            `gorm:"primary_key" json:"-"`
	CreatedAt      time.Time       `json:"-"`
	Chain          string          `json:"chain"`
	Tx             string          `json:"tx" sql:"size:100"`
	Address        string          `json:"address" sql:"size:100"`
	Target         string          `json:"target" sql:"size:100"`
	Currency       string          `json:"currency"`
	BlockNum       int             `json:"block_num"`
	Amount         decimal.Decimal `json:"amount" sql:"type:decimal(40,0);" `
	BlockTimestamp int             `json:"block_timestamp"`
	Deposit        string          `json:"deposit"`
}

func AddRedeemRecord(chain, tx, address, target, currency string, amount decimal.Decimal, blockNum, blockTimestamp int, deposit string) error {
	db := util.DB
	query := db.Create(&RedeemRecord{
		Chain: chain, Tx: tx, Address: address, Target: target, Amount: amount, Currency: currency, BlockNum: blockNum,
		BlockTimestamp: blockTimestamp, Deposit: deposit,
	})
	return query.Error
}

func RedeemList(address string) (list []RedeemRecord) {
	db := util.DB
	db.Model(list).Where("address = ?", address).Order("block_num desc").Find(&list)
	return
}
