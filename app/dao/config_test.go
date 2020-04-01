package dao

import (
	"fmt"
	"testing"
	"time"
)

func TestSetConfig(t *testing.T) {
	sh, _ := time.LoadLocation("Asia/Shanghai")
	expire := "2020-04-01 16:00:00"
	ex, err := time.Parse("2006-01-02 15:04:05", expire)
	now := time.Now().In(sh)
	if ex.Sub(now) < 10*time.Minute {
		fmt.Println(ex)
	}
	fmt.Println(err)

}
