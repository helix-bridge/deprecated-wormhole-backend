package db

import (
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type EthereumLockRecord struct {
	ID             uint            `gorm:"primary_key" json:"-"`
	CreatedAt      time.Time       `json:"-"`
	Chain          string          `json:"chain"`
	Tx             string          `json:"tx" sql:"size:100"`
    Token          string          `json:"token", sql:"size:100"`
	From           string          `json:"from" sql:"size:100"`
	Receiver       string          `json:"receiver" sql:"size:100"`
	BlockNum       int             `json:"block_num"`
	Amount         decimal.Decimal `json:"amount" sql:"type:decimal(40,0);" `
	BlockTimestamp int             `json:"block_timestamp"`
	DarwiniaTx     string          `json:"darwinia_tx"`
	IsRelayed      bool            `json:"is_relayed" gorm:"-"`
}

func AddEthereumLockRecord(chain, tx, token, from, receiver, currency string, amount decimal.Decimal, blockNum, blockTimestamp int) error {
	db := util.DB
	query := db.Create(&EthereumLockRecord{
        Chain: chain, Tx: tx, Token: token, From: from, Receiver: receiver, Amount: amount, BlockNum: blockNum,
		BlockTimestamp: blockTimestamp,
	})
	return query.Error
}

func EthereumLockList(from string) (list []RedeemRecord) {
	db := util.DB
	db.Model(list).Where("from = ?", from).Order("block_num desc").Find(&list)
	for index, value := range list {
		list[index].IsRelayed = GetRelayBestBlockNum() >= uint64(value.BlockNum)
	}
	return
}

func UpdateEthereumLockRecord(tx, darwiniaTx string) {
	util.DB.Model(RedeemRecord{}).Where("tx = ?", tx).Update(RedeemRecord{DarwiniaTx: darwiniaTx})
}

