package db

import (
	"github.com/darwinia-network/link/util"
	"time"
)

type Subscriber struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `gorm:"type:varchar(64);" json:"email" `
}

func CreateSubscribe(email string) error {
	db := util.DB
	query := db.Create(&Subscriber{Email: email})
	return query.Error
}
