package dao

import (
	"huage.tech/mini/app/bean"
)

func GetConfig(key string) (data string) {
	var c bean.Config
	db.Where("key=?", key).First(&c)
	data = c.Data
	return
}

func SetConfig(key, data string) (c bean.Config, err error) {
	err = db.Where("key=? ", key).First(&c).Error
	c.Key = key
	c.Data = data
	if c.ID == 0 {
		err = db.Create(&c).Error
	} else {
		err = db.Model(&c).Update(c).Error
	}
	return
}
