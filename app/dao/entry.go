package dao

import (
	"huage.tech/mini/app/bean"
)

func Entries(roleId, tenantId int64) (result []*bean.EntryTree, err error) {
	//超级管理员
	if roleId == 1 && tenantId == 1 {
		err = db.Model(&bean.EntryTree{}).
			Where("type=?", 1).
			Where("super=?", 1).
			Order("sort").
			Find(&result).Error
	} else {
		//普通管理员
		//TODO
		//从tenant表中获取role_admin,如果==roleId,
		err = db.Model(&bean.EntryTree{}).
			Where("type=?", 1).
			Where("super=?", 0).
			Order("sort").
			Find(&result).Error
		//否则
		//TODO
		err = db.Raw("select e.* from sys_entry e right join sys_auth a"+
			" on e.id = a.entry_id and a.role_id=? where e.type=1 where e.super =0 order by e.sort", roleId).Scan(&result).Error
	}

	return
}

func FindEntry(method, href string) (entry bean.Entry, err error) {
	err = db.Where("method=? and href=?", method, href).First(&entry).Error
	return
}
