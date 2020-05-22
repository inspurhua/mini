package dao

import (
	"fmt"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
)

func MaterialList(tenantId int64, name string, typeId int64, offset, limit int64) (r []bean.MaterialWithFile, count int, err error) {
	countSql := `select count(id) as count from sys_material where tenant_id=? and name like ? `
	countParam := []interface{}{tenantId, "%" + name + "%"}

	retSql := "select a.*,f.name as File,f.url from " + config.Prefix + "material a left join " +
		config.Prefix +
		"file f on a.process_file=f.id where a.tenant_id=? and a.name like ? "
	retParam := []interface{}{tenantId, "%" + name + "%"}
	if typeId > 0 {
		countSql += ` and type_id = ? `
		countParam = append(countParam, typeId)
		retSql += ` and type_id = ? `
		retParam = append(retParam, typeId)
	}
	fmt.Println(countSql)
	fmt.Println(countParam)
	err = db.Raw(countSql, countParam...).Row().Scan(&count)
	if err != nil {
		return
	}
	retSql += " offset ? limit ?"
	retParam = append(retParam, offset, limit)
	err = db.New().Raw(retSql, retParam...).Scan(&r).Error

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
