package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func TenantCreate(c *gin.Context) {
	var err error
	var form bean.Tenant
	RoleId, _ := c.MustGet("Tenant_ID").(int64)
	TenantId, _ := c.MustGet("Tenant_ID").(int64)

	if !(RoleId == 1 && TenantId == 1) {
		util.AbortNewResultErrorOfServer(c, errors.New("当前角色不支持此操作"))
		return
	}

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, err := dao.TenantCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func TenantList(c *gin.Context) {
	RoleId, _ := c.MustGet("Tenant_ID").(int64)
	TenantId, _ := c.MustGet("Tenant_ID").(int64)

	if !(RoleId == 1 && TenantId == 1) {
		util.AbortNewResultErrorOfServer(c, errors.New("当前角色不支持此操作"))
		return
	}

	r, err := dao.TenantList(TenantId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}

func TenantRead(c *gin.Context) {
	RoleId, _ := c.MustGet("Tenant_ID").(int64)
	TokenTenantId, _ := c.MustGet("Tenant_ID").(int64)

	if !(RoleId == 1 && TokenTenantId == 1) {
		util.AbortNewResultErrorOfServer(c, errors.New("当前角色不支持此操作"))
		return
	}

	TenantId := c.Param("id")
	id, err := strconv.ParseInt(TenantId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	r, err := dao.TenantRead(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func TenantUpdate(c *gin.Context) {
	var err error
	var form bean.Tenant

	RoleId, _ := c.MustGet("Tenant_ID").(int64)
	TokenTenantId, _ := c.MustGet("Tenant_ID").(int64)

	if !(RoleId == 1 && TokenTenantId == 1) {
		util.AbortNewResultErrorOfServer(c, errors.New("当前角色不支持此操作"))
		return
	}

	TenantId := c.Param("id")
	id, err := strconv.ParseInt(TenantId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	form.ID = id
	r, err := dao.TenantUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
