package bean

import "huage.tech/mini/app/config"

type Auth struct {
	ID      int64 `json:"id" gorm:"primary_key"`
	RoleId  int64 `json:"role_id"`
	EntryId int64 `json:"entry_id"`
}

func (Auth) TableName() string {
	return config.Prefix + "auth"
}
