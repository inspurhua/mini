package dao

import (
	"fmt"
	"testing"
)

func TestSetConfig(t *testing.T) {
	c, err := SetConfig("aa", "bb")
	fmt.Println(c, err)
	d := GetConfig("aa")
	if d != "bb" {
		t.Fail()
	}
	fmt.Println(d, err)
}
