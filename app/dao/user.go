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

func UserList(account string, roleId, OrgId int, offset, limit int64) (r []bean.User, count int, err error) {
	db := db.Model(&bean.User{})
	if len(account) > 0 {
		db = db.Where("account like ? ", "%"+account+"%")
	}
	if roleId > 0 {
		db = db.Where("role_id = ?", roleId)
	}
	if OrgId > 0 {
		db = db.Where("org_id = ?", OrgId)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&r).Error
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

func ChangePassword(old, new, new1 string) (err error) {

}
