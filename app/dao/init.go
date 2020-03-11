package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"huage.tech/mini/app/config"
	"log"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open(config.DbType, config.DbUrl)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	db.Close()
}
