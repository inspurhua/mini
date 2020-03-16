package dao

import (
	"huage.tech/mini/app/bean"
)

func FileCreate(e bean.File) (result bean.File, err error) {
	result = e
	err = db.Create(&result).Error
	return
}
func FileRead(id int64) (result bean.File, err error) {
	err = db.Where("id=?", id).First(&result).Error
	return
}
