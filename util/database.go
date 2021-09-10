package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	DB *gorm.DB
)

type GormDB struct {
	*gorm.DB
	gdbDone bool
}

func init() {
	dbPath := GetEnv("MYSQL_HOST", "127.0.0.1")
	dbUser := GetEnv("MYSQL_USER", "root")
	dbPass := GetEnv("MYSQL_PASS", "")
	dbName := GetEnv("MYSQL_DB", "darwinia-dapp")
	DB = initMysql(dbPath, dbUser, dbPass, dbName)
	DB.LogMode(false)
}

func initMysql(host, user, pass, db string) *gorm.DB {
	tdb, err := gorm.Open("mysql", user+":"+pass+"@tcp("+host+")/"+db+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	tdb.DB().SetMaxIdleConns(10)
	tdb.DB().SetMaxOpenConns(100)
	tdb.DB().SetConnMaxLifetime(5 * time.Minute)
	return tdb
}

func CloseDB() {
	if DB != nil {
		_ = DB.Close()
	}

}

func DbBegin() *GormDB {
	txn := DB.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

func (c *GormDB) DbCommit() {
	if c.gdbDone {
		return
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		fmt.Println("Fatal error DbCommit", err)
	}
}

func (c *GormDB) DbRollback() {
	if c.gdbDone {
		return
	}
	tx := c.Rollback()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		fmt.Println("Fatal error DbRollback", err)
	}
}

func install(host, user, pass, dbName string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, ""))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	q := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET = `utf8mb4`", dbName)
	_, err = db.Exec(q)
}
