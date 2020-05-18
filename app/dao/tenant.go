package dao

import "huage.tech/mini/app/bean"

func TenantList() (r []bean.Tenant, err error) {
	err = db.Find(&r).Error
	return
}

func TenantCreate(Tenant bean.Tenant) (result bean.Tenant, err error) {
	tx := db.Begin()
	result = Tenant
	err = tx.Create(&result).Error
	if err != nil {
		tx.Rollback()
		return
	}
	r := bean.Role{
		Name:     result.Name + "管理员",
		Status:   1,
		TenantId: result.ID,
	}
	err = tx.Create(&r).Error
	if err != nil {
		tx.Rollback()
		return
	}
	o := bean.Org{
		Name:     result.Name,
		Code:     "100",
		PId:      0,
		Sort:     0,
		TenantId: result.ID,
	}
	err = tx.Create(&o).Error
	if err != nil {
		tx.Rollback()
		return
	}
	result.RoleAdmin = r.ID
	result.RootOrgId = o.ID
	result.RootOrgCode = o.Code
	err = tx.Model(&result).Update(result).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func TenantRead(id int64) (result bean.Tenant, err error) {
	err = db.Where("id=?", id).First(&result).Error
	return
}

func TenantUpdate(Tenant bean.Tenant) (result bean.Tenant, err error) {
	err = db.Model(&result).Update(Tenant).Error
	return
}
