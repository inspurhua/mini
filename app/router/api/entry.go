package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/config"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
)

//func EntryCreate(c *gin.Context) {
//	var err error
//	var form bean.Entry
//
//	err = c.ShouldBind(&form)
//	if err != nil {
//		util.AbortNewResultErrorOfClient(c, err)
//		return
//	}
//
//	r, err := dao.EntryCreate(form)
//	if err != nil || r.ID == 0 {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//
//	c.JSON(200, util.NewResultOKofWrite(r, 1))
//}
//
//func EntryList(c *gin.Context) {
//	r, err := dao.EntryList()
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//	c.JSON(200, util.NewResultOKofRead(r, len(r)))
//	return
//}
//
//func EntryDelete(c *gin.Context) {
//	EntryId := c.Param("id")
//	id, err := strconv.ParseInt(EntryId, 10, 64)
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//	err = dao.EntryDelete(id)
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//	c.JSON(200, util.NewResultOKofWrite(nil, 1))
//}
//
//func EntryRead(c *gin.Context) {
//	EntryId := c.Param("id")
//	id, err := strconv.ParseInt(EntryId, 10, 64)
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//	r, err := dao.EntryRead(id)
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//	c.JSON(200, util.NewResultOKofRead(r, 1))
//}
//
//func EntryUpdate(c *gin.Context) {
//	var err error
//	var form bean.Entry
//	EntryId := c.Param("id")
//	id, err := strconv.ParseInt(EntryId, 10, 64)
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//
//	err = c.ShouldBind(&form)
//	if err != nil {
//		util.AbortNewResultErrorOfClient(c, err)
//		return
//	}
//	form.ID = id
//	r, err := dao.EntryUpdate(form)
//	if err != nil {
//		util.AbortNewResultErrorOfServer(c, err)
//		return
//	}
//
//	c.JSON(200, util.NewResultOKofWrite(r, 1))
//}

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
