package bean

import (
	"huage.tech/mini/app/config"
)

type Config struct {
	ID   int64  `json:"id" gorm:"primary_key"`
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (Config) TableName() string {
	return config.Prefix + "config"
}
