package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
)

func LogList(c *gin.Context) {
	method := c.DefaultQuery("method", "GET")
	uri := c.DefaultQuery("uri", "")
	pag := c.DefaultQuery("page", "1")
	lim := c.DefaultQuery("limit", "20")
	offset, limit, err := util.PageLimit(pag, lim)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, count, err := dao.LogList(method, uri, offset, limit)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, count))
	return
}
