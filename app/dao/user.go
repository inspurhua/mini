package dao

import (
	"huage.tech/mini/app/config"
	"time"
)

type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
	OrgId    int64  `json:"org_id"`
	Status   int    `json:"status"`

	RealName string `json:"real_name"`
	Code     string `json:"code"`
	Position string `json:"position"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Gender   int    `json:"gender"`
	State    int    `json:"state"`

	Note     string     `json:"note"`
	OpenID   string     `json:"open_id"`
	CreateAt time.Time  `json:"create_at"`
	UpdateAt time.Time  `json:"update_at"`
	DeleteAt *time.Time `json:"delete_at"`
}

func (User) TableName() string {
	return config.Prefix + "user"
}
func Login(account, password string) (u User, err error) {
	err = db.Model(&User{}).Where("account=$1 and password=md5($2) and status=1", account,
		config.JwtSecret+password).First(&u).Error
	return
}
