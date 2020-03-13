package dao

import (
	"huage.tech/mini/app/bean"
)

func EntryList() (e []*bean.Entry, err error) {
	err = db.Find(&e).Error
	return
}

func EntryCreate(e *bean.Entry) (result *bean.Entry, err error) {
	result = e
	err = db.Create(result).Error
	return
}

func EntryDelete(id int64) (err error) {
	err = db.Where("id=?", id).Delete(&bean.Entry{}).Error
	return
}

func EntryRead(id int64) (result *bean.Entry, err error) {
	err = db.Where("id=?", id).First(result).Error
	return
}

func EntryUpdate(e *bean.Entry) (result *bean.Entry, err error) {
	err = db.Model(result).Update(e).Error
	return
}

func Entries() (result []*bean.EntryTree, err error) {
	err = db.Model(&bean.EntryTree{}).
		Where("type=?", 1).
		Order("sort").
		Find(&result).Error
	return
}

func FindEntry(method, url string) (entry *bean.Entry, err error) {
	err = db.Where("method=? and url=?", method, url).First(entry).Error
	return
}
