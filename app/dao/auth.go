package dao

import "huage.tech/mini/app/bean"

func FindAuth(RoleId, EntryId int64) (auth *bean.Auth, err error) {
	err = db.Where("role_id=? and entry_id=?", RoleId, EntryId).First(auth).Error
	return
}
