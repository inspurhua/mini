package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func ColsData(c *gin.Context) {
	tenantId, _ := c.MustGet("TENANT_ID").(int64)
	_keyId := c.Param("key_id")
	keyId, err := strconv.ParseInt(_keyId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	r, err := dao.ColsData(keyId, tenantId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}
func TmplData(c *gin.Context) {
	tenantId, _ := c.MustGet("TENANT_ID").(int64)
	_keyId := c.Param("key_id")
	keyId, err := strconv.ParseInt(_keyId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	r, err := dao.TmplData(keyId, tenantId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.String(200, r)
	return
}

func DataList(c *gin.Context) {
	_keyId := c.Param("key_id")

	keyId, err := strconv.ParseInt(_keyId, 10, 64)

	if err != nil {
		keyId = 0
	}

	tenantId, _ := c.MustGet("TENANT_ID").(int64)

	pag := c.DefaultQuery("page", "1")
	lim := c.DefaultQuery("limit", "20")
	offset, limit, err := util.PageLimit(pag, lim)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, count, err := dao.DataList(keyId, tenantId, offset, limit)

	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, count))
	return
}
func DataCreate(c *gin.Context) {
	var err error
	var form bean.QualityData

	_keyId := c.Param("key_id")
	keyId, err := strconv.ParseInt(_keyId, 10, 64)
	if err != nil || keyId == 0 {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	err = c.ShouldBind(&form)
	form.KeyId = keyId
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId

	uId, _ := c.MustGet("UID").(int64)
	form.CreateBy = uId
	r, err := dao.DataCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func DataDelete(c *gin.Context) {
	_keyId := c.Param("key_id")
	keyId, err := strconv.ParseInt(_keyId, 10, 64)
	if err != nil || keyId == 0 {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	_Id := c.Param("id")
	id, err := strconv.ParseInt(_Id, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)

	err = dao.DataDelete(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func DataRead(c *gin.Context) {

	_keyId := c.Param("key_id")
	keyId, err := strconv.ParseInt(_keyId, 10, 64)
	if err != nil || keyId == 0 {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	_Id := c.Param("id")
	id, err := strconv.ParseInt(_Id, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.DataRead(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func DataUpdate(c *gin.Context) {
	var err error
	var form bean.QualityData
	_Id := c.Param("id")
	id, err := strconv.ParseInt(_Id, 10, 64)
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
	r, err := dao.DataUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
