package dao

import (
	"huage.tech/mini/app/bean"
)

func MaterialTypeList(tenantId int64) (e []bean.MaterialType, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Order("sort").
		Find(&e).Error

	return
}

func MaterialTypeCreate(e bean.MaterialType) (result bean.MaterialType, err error) {
	result = e
	err = db.Create(&result).Error
	return
}

func MaterialTypeDelete(tenantId, id int64) (err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).Delete(&bean.MaterialType{}).Error
	return
}

func MaterialTypeRead(tenantId, id int64) (result bean.MaterialType, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).First(&result).Error
	return
}

func MaterialTypeUpdate(e bean.MaterialType) (result bean.MaterialType, err error) {
	err = db.Model(&result).Update(e).Error
	return
}

func MaterialTypeTree(tenantId int64, code string) (result []*bean.MaterialTypeTree, err error) {
	err = db.Model(&bean.MaterialTypeTree{}).
		Where("tenant_id=?", tenantId).
		Where("code like ?", code+"%").
		Order("sort").
		Find(&result).Error
	return
}

func MaterialTypeHasChild(id int64) (has bool, err error) {
	count := 0
	err = db.Model(&bean.MaterialType{}).Where("pid=?", id).Count(&count).Error
	has = count > 0
	return
}
