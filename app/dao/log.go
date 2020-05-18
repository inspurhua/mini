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
	TenantId int64      `json:"tenant_id"`
}

func LogList(date, account, method, uri string, offset, limit int64) (e []bean.LogResult, count int, err error) {
	dd := strings.Split(date, " - ")

	row := db.Raw("select count(a.id) from "+
		config.Prefix+"log a left join "+config.Prefix+"user b "+
		" on a.user_id = b.id where b.account like ? and a.uri like ? and a.method=? "+
		" and a.create_at >= ? and a.create_at <= ?",
		"%"+account+"%",
		"%"+uri+"%",
		method, dd[0], dd[1],
	).Row()
	row.Scan(&count)

	if err != nil {
		return
	}

	err = db.Raw("select b.account,a.id,a.method,a.uri,a.create_at from "+
		config.Prefix+"log a left join "+config.Prefix+"user b "+
		" on a.user_id = b.id where b.account like ? and a.uri like ? and a.method=?"+
		" and a.create_at >= ? and a.create_at <= ?  order by id desc offset ? limit ?",
		"%"+account+"%",
		"%"+uri+"%",
		method, dd[0], dd[1], offset, limit,
	).Scan(&e).Error

	return

}

func LogCreate(e bean.Log) (result bean.Log, err error) {
	result = e
	err = db.Create(&result).Error
	return
}
