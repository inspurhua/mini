package dao

import "huage.tech/mini/app/bean"
//TODO 实现锁有数据库方法,参考db.sql的解释
func TenantList(TenantId int64) (r []bean.Tenant, err error) {
	if TenantId != 1 {
		err = db.Where("id != ?", 1).Find(&r).Error
		return
	}
	err = db.Find(&r).Error
	return
}

func TenantCreate(Tenant bean.Tenant) (result bean.Tenant, err error) {
	result = Tenant
	err = db.Create(&result).Error
	return
}

func TenantDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&bean.Tenant{}).Error
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
