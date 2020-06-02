package dao

import (
	"huage.tech/mini/app/bean"
)

func Test() (result []*bean.Auth, err error) {
	role := bean.Role{
		ID: 8,
	}
	err = db.Model(&role).Related(&result).Error
	return
}
