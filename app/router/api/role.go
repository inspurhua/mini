package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

type RoleForm struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Status int    `form:"status" json:"status" binding:"oneof=0 1"`
}

func RoleCreate(c *gin.Context) {
	var err error
	var req RoleForm

	err = c.ShouldBind(&req)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, err := dao.RoleCreate(req.Name, req.Status)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func RoleList(c *gin.Context) {
	r, err := dao.RoleList()
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
	err = dao.RoleDelete(id)
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
	r, err := dao.RoleRead(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func RoleUpdate(c *gin.Context) {
	var err error
	var req RoleForm
	roleId := c.Param("id")
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	err = c.ShouldBind(&req)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, err := dao.RoleUpdate(&dao.Role{
		ID:     id,
		Name:   req.Name,
		Status: req.Status,
	})
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
