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
		)
		db.Model(RingBurnRecord{}).AddUniqueIndex("tx", "tx")
		db.Model(RingBurnRecord{}).AddIndex("address", "address")
		db.Model(RedeemRecord{}).AddUniqueIndex("tx", "tx")
		db.Model(RedeemRecord{}).AddIndex("address", "address")
		db.Model(Subscriber{}).AddUniqueIndex("email", "email")
		db.Model(DarwiniaBackingLock{}).AddIndex("target", "target")
		db.Model(TokenRegisterRecord{}).AddIndex("source", "source")
		// TODO replaced by sender
		db.Model(TokenBurnRecord{}).AddIndex("recipient", "recipient")
	}
}
