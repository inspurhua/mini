package dao

import (
	"huage.tech/mini/app/bean"
)

func LogList() (e []bean.Log, err error) {
	err = db.Order("create_at desc").Find(&e).Error
	return
}

func LogCreate(e bean.Log) (result bean.Log, err error) {
	result = e
	err = db.Create(&result).Error
	return
}
