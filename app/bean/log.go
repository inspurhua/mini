package bean

import "huage.tech/mini/app/config"

type Log struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	UserId int64  `json:"user_id"`
	Method string `json:"method"`
	Uri    string `json:"uri"`
	Data   string `json:"data"`
	Ip     string `json:"data"`
	Ua     string `json:"data"`
}

func (Log) TableName() string {
	return config.Prefix + "log"
}
