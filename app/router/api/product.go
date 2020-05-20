package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func ProductCreate(c *gin.Context) {
	var err error
	var form bean.Product

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	r, err := dao.ProductCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func ProductList(c *gin.Context) {
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	name := c.DefaultQuery("name", "")
	pag := c.DefaultQuery("page", "1")
	lim := c.DefaultQuery("limit", "20")
	offset, limit, err := util.PageLimit(pag, lim)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	r, err := dao.ProductList(TenantId, name, offset, limit)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}

func ProductDelete(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
 
	err = dao.ProductDelete(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func ProductRead(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.ProductRead(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func ProductUpdate(c *gin.Context) {
	var err error
	var form bean.Product
	productId := c.Param("id")
	id, err := strconv.ParseInt(productId, 10, 64)
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
	r, err := dao.ProductUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
