package dao

import "huage.tech/mini/app/bean"

func RoleList(tenantId, roleId int64) (r []bean.Role, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id != ?", roleId).Find(&r).Error
	return
}

func RoleCreate(role bean.Role) (result bean.Role, err error) {
	result = role
	err = db.Create(&result).Error
	return
}

func RoleDelete(tenantId, id int64) (err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).Delete(&bean.Role{}).Error
	return
}

func RoleRead(tenantId, id int64) (result bean.Role, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).First(&result).Error
	return
}

func RoleUpdate(role bean.Role) (result bean.Role, err error) {
	err = db.Model(&result).Update(role).Error
	return
}
