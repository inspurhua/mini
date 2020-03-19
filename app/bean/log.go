package bean

import (
	"huage.tech/mini/app/config"
	"time"
)

type Log struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	UserId   int64     `json:"user_id"`
	Method   string    `json:"method"`
	Uri      string    `json:"uri"`
	Data     string    `json:"data"`
	Ip       string    `json:"data"`
	Ua       string    `json:"data"`
	CreateAt time.Time `json:"create_at"`
}

func (Log) TableName() string {
	return config.Prefix + "log"
}

type LogResult struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	Account  string    `json:"account"`
	Method   string    `json:"method"`
	Uri      string    `json:"uri"`
	CreateAt time.Time `json:"create_at"`
}
