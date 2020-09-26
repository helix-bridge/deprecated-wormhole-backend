package db

import (
	"github.com/darwinia-network/link/util"
)

func init() {
	db := util.DB
	if db != nil {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
			RingBurnRecord{},
			Subscriber{},
		)
		db.Model(RingBurnRecord{}).AddUniqueIndex("tx", "tx")
		db.Model(RingBurnRecord{}).AddIndex("address", "address")
		db.Model(Subscriber{}).AddUniqueIndex("email", "email")
	}
}
