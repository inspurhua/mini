package dao

import "huage.tech/mini/app/config"

type Role struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Status int    `json:"status" default:1`
}

func (Role) TableName() string {
	return config.Prefix + "role"
}

func RoleList() (r []Role, err error) {
	err = db.Find(&r).Error
	return
}

func RoleCreate(Name string, Status int) (r Role, err error) {
	r = Role{
		Name:   Name,
		Status: Status,
	}
	err = db.Model(&Role{}).Create(&r).Error
	return
}

func RoleDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&Role{}).Error
	return
}

func RoleRead(id int64) (r Role, err error) {
	err = db.Where("id=?", id).First(&r).Error
	return
}

func RoleUpdate(role *Role) (r Role, err error) {
	err = db.Model(&r).Update(role).Error
	//err = db.Model(&r).Where("id=?", id).Updates(m).Error
	//db.Where("id=?", id).First(&r)
	return
}
