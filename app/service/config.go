package service

import (
	"huage.tech/mini/app/bean"
)

var conf = make(map[string]string)

func GetConfig(key string) (data string) {
	if value, ok := conf[key]; ok {
		data = value
	}
	return
}

func SetConfig(key, data string) (c bean.Config, err error) {
	conf[key] = data
	c = bean.Config{
		ID:   101,
		Key:  key,
		Data: data,
	}
	err = nil
	return
}
