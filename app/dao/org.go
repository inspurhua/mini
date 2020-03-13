package dao

import (
	"huage.tech/mini/app/bean"
)

func OrgList() (e []bean.Org, err error) {
	err = db.Order("sort").Find(&e).Error
	return
}

func OrgCreate(e bean.Org) (result bean.Org, err error) {
	result = e
	err = db.Create(&result).Error
	return
}

func OrgDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&bean.Org{}).Error
	return
}

func OrgRead(id int64) (result bean.Org, err error) {
	err = db.Where("id=?", id).First(&result).Error
	return
}

func OrgUpdate(e bean.Org) (result bean.Org, err error) {
	err = db.Model(&result).Update(e).Error
	return
}

func OrgTree() (result []*bean.OrgTree, err error) {
	err = db.Model(&bean.OrgTree{}).
		Order("sort").
		Find(&result).Error
	return
}

func OrgHasChild(id int64) (has bool, err error) {
	count := 0
	err = db.Model(&bean.Org{}).Where("pid=?", id).Count(&count).Error
	has = count > 0
	return
}
