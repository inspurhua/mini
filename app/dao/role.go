package dao

import "huage.tech/mini/app/bean"

func RoleList() (r []*bean.Role, err error) {
	err = db.Find(&r).Error
	return
}

func RoleCreate(role *bean.Role) (result *bean.Role, err error) {
	result = role
	err = db.Create(result).Error
	return
}

func RoleDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&bean.Role{}).Error
	return
}

func RoleRead(id int64) (result *bean.Role, err error) {
	err = db.Where("id=?", id).First(result).Error
	return
}

func RoleUpdate(role *bean.Role) (result *bean.Role, err error) {
	err = db.Model(result).Update(role).Error
	return
}
