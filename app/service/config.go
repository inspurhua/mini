package service

import (
	"huage.tech/mini/app/bean"
	"sync"
)

type Config struct {
	Data map[string]string
	Lock sync.RWMutex
}

var conf = Config{
	Data: make(map[string]string),
}

func GetConfig(key string) (data string) {
	if value, ok := conf.Data[key]; ok {
		data = value
	}
	return
}

func SetConfig(key, data string) (c bean.Config, err error) {
	conf.Lock.Lock()
	defer conf.Lock.Unlock()
	conf.Data[key] = data
	c = bean.Config{
		ID:   101,
		Key:  key,
		Data: data,
	}
	err = nil
	return
}
