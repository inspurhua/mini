package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func MaterialCreate(c *gin.Context) {
	var err error
	var form bean.Material

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	r, err := dao.MaterialCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func MaterialList(c *gin.Context) {
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	name := c.DefaultQuery("name", "")
	pag := c.DefaultQuery("page", "1")
	lim := c.DefaultQuery("limit", "20")
	offset, limit, err := util.PageLimit(pag, lim)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	r, count, err := dao.MaterialList(TenantId, name, offset, limit)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, count))
	return
}

func MaterialDelete(c *gin.Context) {
	MaterialId := c.Param("id")
	id, err := strconv.ParseInt(MaterialId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)

	err = dao.MaterialDelete(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func MaterialRead(c *gin.Context) {
	MaterialId := c.Param("id")
	id, err := strconv.ParseInt(MaterialId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.MaterialRead(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func MaterialUpdate(c *gin.Context) {
	var err error
	var form bean.Material
	MaterialId := c.Param("id")
	id, err := strconv.ParseInt(MaterialId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	err = c.ShouldBind(&form)
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	form.ID = id
	r, err := dao.MaterialUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
