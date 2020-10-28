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
	DarwiniaTx     string          `json:"darwinia_tx"`
	IsRelayed      bool            `json:"is_relayed" gorm:"-"`
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
	for index, value := range list {
		list[index].IsRelayed = GetRelayBestBlockNum() >= uint64(value.BlockNum)
	}
	return
}

func UpdateRedeem(tx, darwiniaTx string) {
	util.DB.Model(RedeemRecord{}).Where("tx = ?", tx).Update(RedeemRecord{DarwiniaTx: darwiniaTx})
}

func SetRelayBestBlockNum(blockNum uint64) {
	_ = util.SetCache("RelayBestBlockNum", blockNum, 86400*30)
}

func GetRelayBestBlockNum() uint64 {
	return util.GetCacheUint64("RelayBestBlockNum")
}
