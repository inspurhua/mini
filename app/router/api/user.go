package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func UserCreate(c *gin.Context) {
	var err error
	var form bean.User

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, err := dao.UserCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func UserList(c *gin.Context) {
	account := c.DefaultQuery("account", "")

	roleId, err := strconv.Atoi(c.DefaultQuery("role_id", ""))

	if err != nil {
		roleId = 0
	}

	orgId, err := strconv.Atoi(c.DefaultQuery("org_id", ""))
	if err != nil {
		orgId = 0
	}
	pag := c.DefaultQuery("page", "1")
	lim := c.DefaultQuery("limit", "20")
	offset, limit, err := util.PageLimit(pag, lim)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	r, count, err := dao.UserList(account, roleId, orgId, offset, limit)

	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, count))
	return
}

func UserDelete(c *gin.Context) {
	UserId := c.Param("id")
	id, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	err = dao.UserDelete(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func UserRead(c *gin.Context) {
	UserId := c.Param("id")
	id, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	r, err := dao.UserRead(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func UserUpdate(c *gin.Context) {
	var err error
	var form bean.User
	UserId := c.Param("id")
	id, err := strconv.ParseInt(UserId, 10, 64)
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
	r, err := dao.UserUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}
