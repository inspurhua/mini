package dao

import (
	"fmt"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
)

func FindAuth(RoleId, EntryId int64) (auth bean.Auth, err error) {
	err = db.Where("role_id=? and entry_id=?", RoleId, EntryId).First(&auth).Error
	return
}

func AuthList(tenantId, RoleId int64) (result []*bean.EntryTree, err error) {
	prefix := config.Prefix
	err = db.Raw(
		fmt.Sprintf(
			"select e.*,a.id as auth_id from %ventry e "+
				"left join %vauth a on e.id = a.entry_id and a.role_id=? "+
				"where e.kind in(0,2) and e.tennantid in (?)",
			prefix, prefix), RoleId, []int64{0, tenantId}).Scan(&result).Error

	return
}

func AuthUpdate(RoleId int64, entryIds []int64) (err error) {
	tx := db.Begin()

	if err := tx.Where("role_id=?", RoleId).Delete(&bean.Auth{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	var auth []bean.Auth

	for _, entryId := range entryIds {
		auth = append(auth, bean.Auth{RoleId: RoleId, EntryId: entryId})
	}
	for _, v := range auth {
		if err := tx.Create(&v).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
