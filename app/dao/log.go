package dao

import (
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/config"
	"strings"
	"time"
)

type LogResult struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	Account  string    `json:"account"`
	Method   string    `json:"method"`
	Uri      string    `json:"uri"`
	CreateAt time.Time `json:"create_at"`
	TenantId int64     `json:"tenant_id"`
}

func LogList(tenantId int64, date, account, method, uri string, offset, limit int64) (e []bean.LogResult, count int, err error) {
	dd := strings.Split(date, " - ")

	segTenant := ""
	slCount := []interface{}{
		"%" + account + "%",
		"%" + uri + "%",
		method, dd[0], dd[1],
	}
	slRet := slCount

	if tenantId > 0 {
		segTenant = " and a.tenant_id=? "
		slCount = append(slCount, tenantId)
		slRet = append(slRet, tenantId, offset, limit)
	} else {
		slRet = append(slRet, offset, limit)
	}

	countSql := "select count(a.id) from " +
		config.Prefix + "log a left join " + config.Prefix + "user b " +
		" on a.user_id = b.id where b.account like ? and a.uri like ? and a.method=? " +
		" and a.create_at >= ? and a.create_at <= ?" + segTenant
	retSql := "select b.account,a.id,a.method,a.uri,a.create_at from " +
		config.Prefix + "log a left join " + config.Prefix + "user b " +
		" on a.user_id = b.id where b.account like ? and a.uri like ? and a.method=?" +
		" and a.create_at >= ? and a.create_at <= ? " + segTenant + " order by id desc offset ? limit ?"

	row := db.Raw(countSql, slCount...).Row()
	row.Scan(&count)

	if err != nil {
		return
	}

	err = db.Raw(retSql, slRet...).Scan(&e).Error

	return

}

func LogCreate(e bean.Log) (result bean.Log, err error) {
	result = e
	err = db.Create(&result).Error
	return
}
