package dao

import (
	"huage.tech/mini/app/bean"
)

func Entries(roleId, tenantId int64) (result []*bean.EntryTree, err error) {
	//超级管理员
	if roleId == 0 && tenantId == 0 {
		err = db.Model(&bean.EntryTree{}).
			Where("type=?", 1).
			Where("kind in (?)", []int{0, 1}).
			Order("sort").
			Find(&result).Error
	} else {
		//从tenant表中获取role_admin,如果==roleId,
		t, _ := TenantRead(tenantId)
		if t.RoleAdmin == roleId {
			//普通管理员
			err = db.Model(&bean.EntryTree{}).
				Where("type=?", 1).
				Where("kind in(?)", []int{0, 2}).
				Order("sort").
				Find(&result).Error
		} else {
			//普通操作员,根据权限表查询
			err = db.Raw("select e.* from sys_entry e right join sys_auth a"+
				" on e.id = a.entry_id and a.role_id=? where e.type=1 where e.kind in (0,2) order by e.sort", roleId).Scan(&result).Error
		}
	}
	return
}

func FindEntry(method, href string) (entry bean.Entry, err error) {
	err = db.Where("method=? and href=?", method, href).First(&entry).Error
	return
}
