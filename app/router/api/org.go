package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func OrgCreate(c *gin.Context) {
	var err error
	var form bean.Org

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	r, err := dao.OrgCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func OrgDelete(c *gin.Context) {
	OrgId := c.Param("id")
	id, err := strconv.ParseInt(OrgId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	has, err := dao.OrgHasChild(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	if has {
		util.AbortNewResultErrorOfServer(c, errors.New("当前部门存在下级部门不能删除"))
		return
	}
	err = dao.OrgDelete(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func OrgUpdate(c *gin.Context) {
	var err error
	var form bean.Org
	OrgId := c.Param("id")
	id, err := strconv.ParseInt(OrgId, 10, 64)
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
	r, err := dao.OrgUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func OrgRead(c *gin.Context) {
	OrgId := c.Param("id")
	id, err := strconv.ParseInt(OrgId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	r, err := dao.OrgRead(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func OrgList(c *gin.Context) {
	r, err := dao.OrgList()
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}

func OrgTree(c *gin.Context) {
	v, err := dao.OrgTree()
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	tree := util.TreeOfOrg(v)
	c.JSON(200, util.NewResultOKofRead(tree, 1))
}