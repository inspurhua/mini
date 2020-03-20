package dao

import (
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
	"time"
)

func Login(account, password string) (u bean.User, err error) {
	err = db.Model(&bean.User{}).Where("account=$1 and password=md5($2) and status=1", account,
		config.JwtSecret+password).First(&u).Error
	return
}

func UserList(account string, roleId int64, orgCode string, offset, limit int64) (r []bean.UserResponse, count int, err error) {
	sql1 := "select count(u.id) from " + config.Prefix + "user u left join " + config.Prefix + "org o on u.org_id=o.id" +
		" where u.account like ? "
	param1 := []interface{}{"%" + account + "%"}

	sql := "select u.id,u.account,u.status,u.role_id,u.org_id,u.update_at,r.name as role,o.name as org from " + config.Prefix + "user u " +
		" left join " + config.Prefix + "role r on u.role_id = r.id " +
		" left join " + config.Prefix + "org o on u.org_id = o.id " +
		" where u.account like ? "
	param := []interface{}{"%" + account + "%"}

	if roleId > 0 {
		sql1 += "and u.role_id = ? "
		param1 = append(param1, roleId)

		sql += "and u.role_id = ? "
		param = append(param, roleId)
	}
	if len(orgCode) > 0 {
		sql1 += " and o.code like ? "
		param1 = append(param1, orgCode+"%")

		sql += " and o.code like ? "
		param = append(param, orgCode+"%")
	}
	err = db.Raw(sql1, param1...).Row().Scan(&count)
	if err != nil {
		return
	}
	sql += " order by u.update_at desc offset ? limit ?"
	param = append(param, offset, limit)
	err = db.New().Raw(sql, param...).Scan(&r).Error

	return
}

func UserCreate(User bean.User) (result bean.User, err error) {
	now := time.Now()
	User.CreateAt = now
	User.UpdateAt = now
	result = User

	err = db.Create(&result).Error
	result.Password =""
	return
}

func UserDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&bean.User{}).Error
	return
}

func UserRead(id int64) (result bean.User, err error) {
	err = db.Where("id=?", id).First(&result).Error
	result.Password =""
	return
}

func UserUpdate(User bean.User) (result bean.User, err error) {
	User.UpdateAt = time.Now()
	err = db.Model(&result).Update(User).Error
	result.Password =""
	return
}
