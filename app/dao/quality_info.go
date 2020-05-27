package dao

import "huage.tech/mini/app/bean"

func QualityInfoList(tenantId, keyId int64) (r []bean.QualityInfo, err error) {
	sql := `
with qc as (
select
	case
		when substring(a.attname from '^(.*)_') = 'key' then 'A条件类'
		else 'B指标类'
	end as type, 
	
	case substring(a.attname from '_(\w)\d+')
when 't' then '文本型'
when 'i' then '整数型'
when 'n' then '小数型'
end as data_type,
a.attname as col_name
from
	pg_class as c
inner join pg_attribute as a on
	a.attrelid = c.oid
	and a.attnum>0
where
	c.relname = 'sys_quality'
	and a.attname != 'key_id'
	and (a.attname like 'key%'
	or a.attname like 'data%') )
select
	qc.*,
	qi.col_title ,
	qi.group_title ,
	qi.refer_text ,
	qi.refer_unit ,
	qi.refer_expr ,
	qi.sort
from
	qc
left join sys_quality_info qi on
	qc.col_name = qi.col_name
	and qi.key_id = ?
	and qi.tenant_id = ? 
order by type,sort,col_name
`
	err = db.Raw(sql, keyId, tenantId).Find(&r).Error
	return
}

func QualityInfoUpdate(tenantId, keyId int64, data []bean.QualityInfo) (err error) {
	tx := db.Begin()
	err = tx.Exec(`delete from sys_quality_info where tenant_id = ? and key_id=?`, tenantId, keyId).Error
	if err != nil {
		tx.Rollback()
		return
	}
	for _, qi := range data {
		qi.TenantId = tenantId
		qi.KeyID = keyId
		err = tx.Create(&qi).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	//检查是否有entry,没有则生成


	tx.Commit()
	return
}
