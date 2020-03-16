package bean

import (
	"huage.tech/mini/app/config"
	"time"
)

type File struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	Name     string    `json:"name" form:"name"`
	SaveName string    `json:"save_name" form:"save_name"`
	SavePath string    `json:"save_path" form:"save_path"`
	Url      string    `json:"url" form:"url"`
	CreateAt time.Time `json:"create_at"`
}

func (File) TableName() string {
	return config.Prefix + "file"
}
