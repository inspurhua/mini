package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"time"
)

func LogList(c *gin.Context) {
	date := c.DefaultQuery("date", "")
	account := c.DefaultQuery("account", "")
	method := c.DefaultQuery("method", "GET")
	uri := c.DefaultQuery("uri", "")
	pag := c.DefaultQuery("page", "1")
	lim := c.DefaultQuery("limit", "20")
	offset, limit, err := util.PageLimit(pag, lim)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	if date == "" {
		now := time.Now()
		bef := now.Add(-1 * 24 * time.Hour)
		date = bef.Format("2006-01-02 15:04:05") + " - " + now.Format("2006-01-02 15:04:05")
	}
	r, count, err := dao.LogList(date, account, method, uri, offset, limit)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, count))
	return
}
