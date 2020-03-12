package bean

import "huage.tech/mini/app/config"

type Role struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	Name   string `form:"name" json:"name" binding:"required"`
	Status int    `form:"status" json:"status" default:1 binding:"oneof=0 1"`
}

func (Role) TableName() string {
	return config.Prefix + "role"
}
