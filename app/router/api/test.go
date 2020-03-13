package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
)

func Test(c *gin.Context) {
	data, err := dao.Test()
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(data, len(data)))

}
