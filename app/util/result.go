package util

import "github.com/gin-gonic/gin"

type Result struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func NewResultOKofRead(Data interface{}, Count int) *Result {
	return NewResultOK("获取成功", Data, Count)
}
func NewResultOKofWrite(Data interface{}, Count int) *Result {
	return NewResultOK("操作成功", Data, Count)
}
func NewResultOK(Msg string, Data interface{}, Count int) *Result {
	return &Result{
		Code:  0,
		Msg:   Msg,
		Data:  Data,
		Count: Count,
	}
}

func NewResultError(Code int, Msg error, Data interface{}, Count int) *Result {
	return &Result{
		Code:  Code,
		Msg:   Msg.Error(),
		Data:  Data,
		Count: Count,
	}
}

func AbortNewResultErrorOfClient(c *gin.Context, err error) {
	c.AbortWithStatusJSON(200, Result{
		Code:  400,
		Msg:   err.Error(),
		Data:  nil,
		Count: 0,
	})
}

func AbortNewResultErrorOfServer(c *gin.Context, err error) {
	c.AbortWithStatusJSON(200, Result{
		Code:  500,
		Msg:   err.Error(),
		Data:  nil,
		Count: 0,
	})
}
