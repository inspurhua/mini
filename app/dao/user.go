package dao

import (
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
)

func Login(account, password string) (u bean.User, err error) {
	err = db.Model(&bean.User{}).Where("account=$1 and password=md5($2) and status=1", account,
		config.JwtSecret+password).First(&u).Error
	return
}

func UserList() (r []bean.User, err error) {
	err = db.Find(&r).Error
	return
}

func UserCreate(User bean.User) (result bean.User, err error) {
	result = User
	err = db.Create(&result).Error
	return
}

func UserDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&bean.User{}).Error
	return
}

func UserRead(id int64) (result bean.User, err error) {
	err = db.Where("id=?", id).First(&result).Error
	return
}

func UserUpdate(User bean.User) (result bean.User, err error) {
	err = db.Model(&result).Update(User).Error
	return
}
