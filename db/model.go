package db

import (
	"github.com/darwinia-network/link/util"
)

func init() {
	db := util.DB
	if db != nil {
		db.Set("gorm:table_options", "ENGINE=InnoDB").
			AutoMigrate()
	}
}
