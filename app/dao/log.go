package dao

import (
	"huage.tech/mini/app/bean"
)

func LogList(method, uri string, offset, limit int64) (e []bean.Log, count int, err error) {
	err = db.Where("method = ? and uri like ?", method, "%"+uri+"%").Count(&count).Error
	if err != nil {
		return
	}
	err = db.Where("method = ? and url like ?", method, "%"+uri+"%").
		Offset(offset).Limit(limit).
		Order("create_at desc").Find(&e).Error
	return
}

func LogCreate(e bean.Log) (result bean.Log, err error) {
	result = e
	err = db.Create(&result).Error
	return
}
