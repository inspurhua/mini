package util

import (
	"crypto/md5"
	"fmt"
	"huage.tech/mini/app/bean"
	"reflect"
	"strconv"
)

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func TreeOfEntry(rows []*bean.EntryTree) (tree []*bean.EntryTree) {
	tmp := make(map[int64]*bean.EntryTree)
	for _, v := range rows {
		tmp[v.ID] = v
	}
	for _, v := range rows {
		if parent, ok := tmp[v.PId]; ok {
			parent.Child = append(parent.Child, v)
		} else {
			tree = append(tree, v)
		}
	}
	return
}

func TreeOfOrg(rows []*bean.OrgTree) (tree []*bean.OrgTree) {
	tmp := make(map[int64]*bean.OrgTree)
	for _, v := range rows {
		tmp[v.ID] = v
	}
	for _, v := range rows {
		if parent, ok := tmp[v.PId]; ok {
			parent.Children = append(parent.Children, v)
		} else {
			tree = append(tree, v)
		}
	}
	return
}

func TreeOfMaterialType(rows []*bean.MaterialTypeTree) (tree []*bean.MaterialTypeTree) {
	tmp := make(map[int64]*bean.MaterialTypeTree)
	for _, v := range rows {
		tmp[v.ID] = v
	}
	for _, v := range rows {
		if parent, ok := tmp[v.PId]; ok {
			parent.Children = append(parent.Children, v)
		} else {
			tree = append(tree, v)
		}
	}
	return
}

func PageLimit(pag, lim string) (offset int64, limit int64, err error) {
	page, err := strconv.ParseInt(pag, 10, 64)
	if err != nil {
		return
	}
	limit, err = strconv.ParseInt(lim, 10, 64)
	if err != nil {
		return
	}

	offset = (page - 1) * limit
	return
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
