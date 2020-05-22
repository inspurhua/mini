package bean

import "huage.tech/mini/app/config"

type QualityInfo struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	KeyID      int64  `json:"-"`
	Type       string `form:"type" json:"type" gorm:"-"`
	ColName    string `form:"col_name" json:"col_name"`
	ColTitle   string `form:"col_title" json:"col_title" binding:"required"`
	GroupTitle string `form:"group_title" json:"group_title"`
	ReferText  string `form:"refer_text" json:"refer_text"`
	ReferExpr  string `form:"refer_expr" json:"-"`
	ReferUnit  string `form:"refer_unit" json:"refer_unit"`
	Sort       int    `json:"sort"  json:"sort"`
	TenantId   int64  `json:"-"`
}

func (QualityInfo) TableName() string {
	return config.Prefix + "quality_info"
}
