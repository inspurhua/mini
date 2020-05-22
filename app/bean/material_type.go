package bean

import "huage.tech/mini/app/config"

type MaterialType struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Name     string `json:"name" form:"name" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
	PId      int64  `json:"pid" form:"pid" gorm:"column:pid" binding:"required"`
	Sort     int    `json:"sort" form:"sort"`
	TenantId int64  `json:"-"`
}

func (MaterialType) TableName() string {
	return config.Prefix + "material"
}

type MaterialTypeTree struct {
	*MaterialType
	Children []*MaterialTypeTree `json:"children" gorm:"-"`
}

func (MaterialTypeTree) TableName() string {
	return config.Prefix + "material"
}
