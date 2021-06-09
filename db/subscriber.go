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

type PloSubscriber struct {
	Email      string `gorm:"type:varchar(64);" json:"email" `
	KsmAddress string `json:"ksm_address"`
}

func CreatePloSubscriber(email, address string) error {
	db := util.DB
	query := db.Create(&PloSubscriber{Email: email, KsmAddress: address})
	return query.Error
}
