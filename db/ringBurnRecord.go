package db

import (
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type RingBurnRecord struct {
	ID        uint            `gorm:"primary_key" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	Chain     string          `json:"chain"`
	Tx        string          `json:"tx" sql:"size:100"`
	Address   string          `json:"address" sql:"size:100"`
	Target    string          `json:"target" sql:"size:100"`
	Currency  string          `json:"currency"`
	Amount    decimal.Decimal `json:"amount" sql:"type:decimal(40,0);" `
}

func AddRingBurnRecord(chain, tx, address, target, currency string, amount decimal.Decimal) error {
	db := util.DB
	query := db.Create(&RingBurnRecord{
		Chain: chain, Tx: tx, Address: address, Target: target, Amount: amount, Currency: currency,
	})
	return query.Error
}
