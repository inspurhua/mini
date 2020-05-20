package dao

import (
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
)

func ProductList(tenantId int64, name string, offset, limit int64) (r []bean.ProductWithFile, count int, err error) {
	err = db.Model(&bean.Product{}).
		Where("tenant_id=?", tenantId).
		Where("name like ?", "%"+name+"%").
		Count(&count).Error
	if err != nil {
		return
	}

	err = db.Raw("select a.*,f.name as File,f.url from "+config.Prefix+"product a left join "+
		config.Prefix+
		"file f on a.process_file=f.id where a.tenant_id=? and a.name like ? offset ? limit ?",
		tenantId, "%"+name+"%", offset, limit).Scan(&r).Error

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

func ProductRead(tenantId, id int64) (result bean.ProductWithFile, err error) {
	err = db.Raw("select a.*,f.name as File,f.url from "+config.Prefix+"product a left join "+
		config.Prefix+
		"file f on a.process_file=f.id where a.tenant_id=? and a.id=?",
		tenantId, id).Scan(&result).Error

	return
}

func ProductUpdate(Product bean.Product) (result bean.Product, err error) {
	err = db.Model(&result).Update(Product).Error
	return
}
