package util

import (
	"huage.tech/mini/app/bean"
	"reflect"
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

func MenuTree(rows []*bean.EntryTree) (tree []*bean.EntryTree) {
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
