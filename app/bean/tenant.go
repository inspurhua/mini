package bean

import "huage.tech/mini/app/config"

type Tenant struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Name        string `form:"name" json:"name" binding:"required"`
	Status      int    `form:"status" json:"status" default:1 binding:"oneof=1 2"`
	RoleAdmin   int64
	RootOrgId   int64
	RootOrgCode string
}

func (Tenant) TableName() string {
	return config.Prefix + "tenant"
}