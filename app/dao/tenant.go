package dao

import "huage.tech/mini/app/bean"

func TenantList() (r []bean.Tenant, err error) {
	err = db.Find(&r).Error
	return
}

func TenantCreate(Tenant bean.Tenant) (result bean.Tenant, err error) {
	result = Tenant
	err = db.Create(&result).Error
	if err != nil {
		return
	}
	r, err := RoleCreate(bean.Role{
		Name:   result.Name + "admin",
		Status: 1,
	})
	if err != nil {
		return
	}
	o, err := OrgCreate(bean.Org{
		Name: result.Name,
		Code: "100",
		PId:  0,
		Sort: 0,
	})
	if err != nil {
		return
	}
	result.RoleAdmin = r.ID
	result.RootOrgId = o.ID
	result.RootOrgCode = o.Code
	result, err = TenantUpdate(result)

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
