package bean

import "huage.tech/mini/app/config"

type Org struct {
	ID   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name" form:"name" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	PId  int64  `json:"pid" form:"pid" gorm:"column:pid" binding:"required"`
	Sort int    `json:"sort" form:"sort"`
}

func (Org) TableName() string {
	return config.Prefix + "org"
}

type OrgTree struct {
	*Org
	Child []*OrgTree `json:"child" gorm:"-"`
}

func (OrgTree) TableName() string {
	return config.Prefix + "org"
}
