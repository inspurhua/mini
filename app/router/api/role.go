package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func RoleCreate(c *gin.Context) {
	var err error
	var form bean.Role

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	r, err := dao.RoleCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func RoleList(c *gin.Context) {
	roleId, _ := c.MustGet("ROLE_ID").(int64)
	TenantId, _ := c.MustGet("TENANT_ID").(int64)

	r, err := dao.RoleList(TenantId, roleId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}

func RoleDelete(c *gin.Context) {
	roleId := c.Param("id")
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	t, err := dao.TenantRead(TenantId)
	if err != nil || t.RoleAdmin == id {
		util.AbortNewResultErrorOfClient(c,
			errors.New(err.Error()+"此角色不能删除"))
		return
	}
	err = dao.RoleDelete(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func RoleRead(c *gin.Context) {
	roleId := c.Param("id")
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.RoleRead(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func RoleUpdate(c *gin.Context) {
	var err error
	var form bean.Role
	roleId := c.Param("id")
	id, err := strconv.ParseInt(roleId, 10, 64)
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
	r, err := dao.RoleUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
