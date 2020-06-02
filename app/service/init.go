package service

import "sync"

var cm *configMap
var once sync.Once

func init() {
	once.Do(func() {
		cm = &configMap{
			data: make(map[string]string, 0),
		}
	})
}

func GetConfigMap() *configMap {
	return cm
}
