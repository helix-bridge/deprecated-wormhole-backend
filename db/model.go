package db

import (
	"github.com/darwinia-network/link/util"
)

func init() {
	db := util.DB
	if db != nil {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
			RingBurnRecord{},
			RedeemRecord{},
			Subscriber{},
			DarwiniaBackingLock{},
			TokenRegisterRecord{},
			TokenBurnRecord{},
			EthereumLockRecord{},
			PloSubscriber{},
		)
		db.Model(RingBurnRecord{}).AddUniqueIndex("tx", "tx")
		db.Model(RingBurnRecord{}).AddIndex("address", "address")
		db.Model(RedeemRecord{}).AddUniqueIndex("tx", "tx")
		db.Model(RedeemRecord{}).AddIndex("address", "address")
		db.Model(Subscriber{}).AddUniqueIndex("email", "email")
		db.Model(PloSubscriber{}).AddUniqueIndex("email", "email")
		db.Model(DarwiniaBackingLock{}).AddIndex("target", "target")
		db.Model(TokenRegisterRecord{}).AddIndex("source", "source")
		db.Model(TokenRegisterRecord{}).AddIndex("tx", "tx")
		db.Model(TokenBurnRecord{}).AddIndex("sender", "sender")
		db.Model(TokenBurnRecord{}).AddIndex("tx", "tx")
		db.Model(EthereumLockRecord{}).AddIndex("sender", "sender")
		db.Model(DarwiniaBackingLock{}).AddIndex("tx", "tx")
		db.Model(DarwiniaBackingLock{}).AddIndex("account_id", "account_id")
	}
}
