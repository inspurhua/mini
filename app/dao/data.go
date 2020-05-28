package dao

import (
	"fmt"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/util"
	"strings"
	"time"
)

func ColsData(keyId, tenantId int64) (r []bean.LayCol, err error) {
	sql := `
select col_name as field, col_title as title
from sys_quality_info
where key_id = ?
  and tenant_id = ?
order by sort
;`
	err = db.Raw(sql, keyId, tenantId).Find(&r).Error
	r = append(r,
		bean.LayCol{
			Field: "qc_date",
			Title: "检测日期",
		}, bean.LayCol{
			Field: "create_at",
			Title: "上传日期",
		}, bean.LayCol{
			Field: "create_by",
			Title: "上传人",
		})
	r = append([]bean.LayCol{
		{
			Field: "batch",
			Title: "批次",
		},
	}, r...)
	//再加上标准的头
	return
}
func TmplData(keyId, tenantId int64) (resp string, err error) {
	sql := `
select col_name as field, col_title as title
from sys_quality_info
where key_id = ?
  and tenant_id = ?
order by sort
;`
	var r []bean.LayCol
	err = db.Raw(sql, keyId, tenantId).Find(&r).Error
	if err != nil {
		return
	}
	r = append([]bean.LayCol{
		{
			Field: "batch",
			Title: "批次",
		}, {
			Field: "qc_date",
			Title: "检测日期",
		},
	}, r...)

	for _, v := range r {
		resp += fmt.Sprintf(`<div class="layui-form layuimini-form">
          <div class="layui-form-item">
              <label class="layui-form-label">%[1]s</label>
              <div class="layui-input-block">
                  <input type="text" name="%[2]s" placeholder="请输入%[1]s"
                         value="{{ d.%[2]s || '' }}"
                         class="layui-input">
              </div>
          </div>`, v.Title, v.Field)
	}

	resp = `<script type="text/html" id="tpl_curd">
  <div class="layuimini-main">
      <div class="layui-form layuimini-form">` + resp + ` <div class="layui-form-item">
              <div class="layui-input-block">
                  <button class="layui-btn" lay-submit lay-filter="saveBtn">保存</button>
              </div>
          </div></div></div></div>`
	return
}
func DataList(keyId, tenantId, offset, limit int64) (result []map[string]interface{}, count int, err error) {
	//api/data/:key_id
	r, err := ColsData(keyId, tenantId)

	if err != nil {
		return
	}
	var cols []string

	for _, x := range r {
		cols = append(cols, x.Field)
	}
	counter := `select count(1) from sys_quality q where key_id=? and tenant_id = ?`
	reporter := `select id,key_id,batch,` + strings.Join(cols, ",") + `,qc_date,create_at from sys_quality q ` +
		` where key_id=? and tenant_id = ? order by id desc offset ? limit ?`
	err = db.Raw(counter, keyId, tenantId).Row().Scan(&count)
	if err != nil {
		return
	}

	rows, err := db.Raw(reporter, keyId, tenantId, offset, limit).Rows()
	if err != nil {
		return
	}
	result, err = util.GetRows(rows)
	return
}
func DataCreate(d bean.QualityData) (result bean.QualityData, err error) {
	d.CreateAt = time.Now()
	result = d
	err = db.Create(&result).Error
	return
}

func DataDelete(tenantId, id int64) (err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).Delete(&bean.QualityData{}).Error
	return
}

func DataRead(tenantId, id int64) (result bean.QualityData, err error) {
	err = db.
		Where("tenant_id=?", tenantId).
		Where("id=?", id).First(&result).Error
	return
}

func DataUpdate(d bean.QualityData) (result bean.QualityData, err error) {
	err = db.Model(&result).Update(d).Error
	return
}
