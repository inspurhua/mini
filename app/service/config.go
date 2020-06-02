package service

import (
	"huage.tech/mini/app/bean"
	"sync"
)

type configMap struct {
	data map[string]string
	lock sync.RWMutex
}

func (c *configMap) GetConfig(key string) (data string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if value, ok := c.data[key]; ok {
		data = value
	}
	return
}

func (c *configMap) SetConfig(key, data string) (ret bean.Config, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = data
	ret = bean.Config{
		ID:   101,
		Key:  key,
		Data: data,
	}
	err = nil
	return
}
