package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/config"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
)

func Entries(c *gin.Context) {
	roleId, _ := c.MustGet("ROLE_ID").(int64)
	tenantId, _ := c.MustGet("TENANT_ID").(int64)
	v, err := dao.Entries(roleId, tenantId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	var tree = util.TreeOfEntry(v)

	c.JSON(200, util.NewResultOKofRead(gin.H{
		"homeInfo": gin.H{
			"title": "首页",
			"href":  "page/welcome-1.html",
		},
		"logoInfo": gin.H{
			"title": config.AppName,
			"image": "images/logo.png",
			"href":  ""},
		"menuInfo": tree}, 1))

}
