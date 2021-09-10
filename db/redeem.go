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

func RedeemList(address string, page, row int, confirmed string) ([]RedeemRecord, int) {
    db := util.DB
    var list []RedeemRecord
    var count int
    switch confirmed {
    case "true":
        db.Model(RedeemRecord{}).Where("darwinia_tx <> ''").Where("address = ?", address).Count(&count)
        db.Where("darwinia_tx <> ''").Where("address = ?", address).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)
    case "false":
        db.Model(RedeemRecord{}).Where("darwinia_tx = ''").Where("address = ?", address).Count(&count)
        db.Where("darwinia_tx = ''").Where("address = ?", address).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)
    default:
        db.Model(RedeemRecord{}).Where("address = ?", address).Count(&count)
        if row > 0 {
            db.Where("address = ?", address).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)
        } else {
            db.Model(list).Where("address = ?", address).Order("block_num desc").Find(&list)
        }
    }
    relay_best_blocknum := GetRelayBestBlockNum()
    for index, value := range list {
        list[index].IsRelayed = relay_best_blocknum >= uint64(value.BlockNum)
    }
    return list, count
}

func UpdateRedeem(tx, darwiniaTx string) {
	util.DB.Model(RedeemRecord{}).Where("tx = ?", tx).Update(RedeemRecord{DarwiniaTx: darwiniaTx})
}

func SetRelayBestBlockNum(blockNum uint64) {
	if blockNum > GetRelayBestBlockNum() {
		_ = util.SetCache("RelayBestBlockNum", blockNum, 86400*30)
	}

}

func GetRelayBestBlockNum() uint64 {
	return util.GetCacheUint64("RelayBestBlockNum")
}

type MappingTradeStat struct {
	TxCount int                    `json:"tx_count"`
	D2E     map[string]interface{} `json:"d2e"`
	E2D     map[string]interface{} `json:"e2d"`
}

func MappingStat() *MappingTradeStat {
	db := util.DB
	type NSum struct {
		N decimal.Decimal
	}
	var n NSum
	pre := decimal.New(1, 18)
	var r MappingTradeStat
	r.E2D = make(map[string]interface{})
	db.Model(RedeemRecord{}).Select("sum(amount) as n").Where("currency = 'ring'").Scan(&n)
	r.E2D["ring"] = n.N.Div(pre)
	db.Model(RedeemRecord{}).Select("sum(amount) as n").Where("currency = 'kton'").Scan(&n)
	r.E2D["kton"] = n.N.Div(pre)
	db.Model(RedeemRecord{}).Select("sum(amount) as n").Where("currency = 'deposit'").Scan(&n)
	r.E2D["deposit"] = n.N.Div(pre)
	var count int
	db.Model(RedeemRecord{}).Count(&count)
	r.E2D["tx_count"] = count
	r.TxCount = count

	r.D2E = make(map[string]interface{})
	db.Model(DarwiniaBackingLock{}).Select("sum(ring_value) as n").Scan(&n)
	r.D2E["ring"] = n.N.Div(decimal.New(1, 9))
	db.Model(DarwiniaBackingLock{}).Select("sum(kton_value) as n").Scan(&n)
	r.D2E["kton"] = n.N.Div(decimal.New(1, 9))
	db.Model(DarwiniaBackingLock{}).Count(&count)
	r.D2E["tx_count"] = count
	r.TxCount += count

	return &r
}
