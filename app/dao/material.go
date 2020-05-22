package dao

import (
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
)

func MaterialList(tenantId int64, name string, offset, limit int64) (r []bean.MaterialWithFile, count int, err error) {
	err = db.Model(&bean.Material{}).
		Where("tenant_id=?", tenantId).
		Where("name like ?", "%"+name+"%").
		Count(&count).Error
	if err != nil {
		return
	}

	err = db.Raw("select a.*,f.name as File,f.url from "+config.Prefix+"Material a left join "+
		config.Prefix+
		"file f on a.process_file=f.id where a.tenant_id=? and a.name like ? offset ? limit ?",
		tenantId, "%"+name+"%", offset, limit).Scan(&r).Error

	return

}

func MaterialCreate(Material bean.Material) (result bean.Material, err error) {
	result = Material
	err = db.Create(&result).Error
	return
}

func MaterialDelete(tenantId, id int64) (err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).
		Delete(&bean.Material{}).Error
	return
}

func MaterialRead(tenantId, id int64) (result bean.MaterialWithFile, err error) {
	err = db.Raw("select a.*,f.name as File,f.url from "+config.Prefix+"Material a left join "+
		config.Prefix+
		"file f on a.process_file=f.id where a.tenant_id=? and a.id=?",
		tenantId, id).Scan(&result).Error

	return
}

func MaterialUpdate(Material bean.Material) (result bean.Material, err error) {
	err = db.Model(&result).Update(Material).Error
	return
}
