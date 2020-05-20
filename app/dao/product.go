package dao

import (
	"huage.tech/mini/app/bean"
)

func ProductList(tenantId int64, name string, offset, limit int64) (r []bean.Product, err error) {
	err = db.Model(&bean.Product{}).
		Where("tenant_id=?", tenantId).
		Where("name like ?", "%"+name+"%").
		Count(&count).Error
	if err != nil {
		return
	}
	err = db.
		Where("tenant_id=?", tenantId).
		Where("name like ?", "%"+name+"%").
		Offset(offset).Limit(limit).
		Find(&r).Error
	return

}

func ProductCreate(Product bean.Product) (result bean.Product, err error) {
	result = Product
	err = db.Create(&result).Error
	return
}

func ProductDelete(tenantId, id int64) (err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).
		Delete(&bean.Product{}).Error
	return
}

func ProductRead(tenantId, id int64) (result bean.Product, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).First(&result).Error

	return
}

func ProductUpdate(Product bean.Product) (result bean.Product, err error) {
	err = db.Model(&result).Update(Product).Error
	return
}
