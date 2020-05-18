package dao

import (
	"huage.tech/mini/app/bean"
)

func OrgList(tenantId int64) (e []bean.Org, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Order("sort").
		Find(&e).Error

	return
}

func OrgCreate(e bean.Org) (result bean.Org, err error) {
	result = e
	err = db.Create(&result).Error
	return
}

func OrgDelete(tenantId, id int64) (err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).Delete(&bean.Org{}).Error
	return
}

func OrgRead(tenantId, id int64) (result bean.Org, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).First(&result).Error
	return
}

func OrgUpdate(e bean.Org) (result bean.Org, err error) {
	err = db.Model(&result).Update(e).Error
	return
}

func OrgTree(tenantId int64, code string) (result []*bean.OrgTree, err error) {
	err = db.Model(&bean.OrgTree{}).
		Where("tenant_id=?", tenantId).
		Where("code like ?", code+"%").
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
