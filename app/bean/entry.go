package bean

import (
	"huage.tech/mini/app/config"
)

type Entry struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Title    string `json:"title" form:"title"`
	PId      int64  `json:"pid" form:"pid" gorm:"column:pid"`
	Type     int    `json:"type" form:"type"`
	Method   string `json:"method" form:"method"`
	Href     string `json:"href" form:"href"`
	Icon     string `json:"icon" form:"icon"`
	Target   string `json:"target" form:"target"`
	Sort     int    `json:"sort" form:"sort"`
	AuthId   int64  `json:"auth_id"`
	TenantId int64  `json:"-"`
}

func (Entry) TableName() string {
	return config.Prefix + "entry"
}

type EntryTree struct {
	*Entry
	Child []*EntryTree `json:"child" gorm:"-"`
}

func (EntryTree) TableName() string {
	return config.Prefix + "entry"
}
