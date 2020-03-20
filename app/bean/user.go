package bean

import (
	"huage.tech/mini/app/config"
	"time"
)

type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	RoleId   int64  `json:"role_id" form:"role_id"`
	OrgId    int64  `json:"org_id" form:"org_id"`
	Status   int    `json:"status" form:"status"`

	RealName string `json:"real_name" form:"real_name"`
	Code     string `json:"code" form:"code"`
	Position string `json:"position" form:"position"`
	Email    string `json:"email" form:"email"`
	Tel      string `json:"tel" form:"tel"`
	Avatar   string `json:"avatar" form:"avatar"`
	Address  string `json:"address" form:"address"`
	Gender   int    `json:"gender" form:"gender"`
	State    int    `json:"state" form:"state"`

	Note     string     `json:"note" form:"note"`
	OpenID   string     `json:"open_id"`
	CreateAt time.Time  `json:"create_at"`
	UpdateAt time.Time  `json:"update_at"`
	DeleteAt *time.Time `json:"delete_at"`
}

func (User) TableName() string {
	return config.Prefix + "user"
}

type UserResponse struct {
	ID       int64     `gorm:"primary_key" json:"id"`
	Account  string    `json:"account" form:"account"`
	Status   int       `json:"status" form:"status"`
	Role     string    `json:"role" form:"role"`
	Org      string    `json:"org" form:"org"`
	RoleId   int64     `json:"role_id" form:"role_id"`
	OrgId    int64     `json:"org_id" form:"org_id"`
	UpdateAt time.Time `json:"update_at"`
}
