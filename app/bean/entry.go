package bean

import "huage.tech/mini/app/config"

type Entry struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	Title  string `json:"title" form:"title"`
	PID    int64  `json:"pid" form:"pid" gorm:"column:pid"`
	Type   int    `json:"type" form:"type"`
	Url    string `json:"url" form:"url"`
	Icon   string `json:"icon" form:"icon"`
	Target string `json:"target" form:"target"`
	Sort   int    `json:"sort" form:"sort"`
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
