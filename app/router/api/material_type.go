package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
	"strings"
)

func MaterialTypeCreate(c *gin.Context) {
	var err error
	var form bean.MaterialType

	err = c.ShouldBind(&form)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}
	err = c.ShouldBind(&form)
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	//check code
	//parent
	parentMaterialType, err := dao.MaterialTypeRead(TenantId, form.PId)
	if !strings.HasPrefix(form.Code, parentMaterialType.Code) || len(parentMaterialType.Code)+3 != len(form.Code) {
		util.AbortNewResultErrorOfClient(c, errors.New("编码必须满足3-3-3...的格式,且以父级编码为前缀"))
		return
	}
	r, err := dao.MaterialTypeCreate(form)
	if err != nil || r.ID == 0 {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func MaterialTypeDelete(c *gin.Context) {
	MaterialTypeId := c.Param("id")
	id, err := strconv.ParseInt(MaterialTypeId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	has, err := dao.MaterialTypeHasChild(id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	if has {
		util.AbortNewResultErrorOfServer(c, errors.New("当前部门存在下级部门不能删除"))
		return
	}

	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	t, err := dao.TenantRead(TenantId)
	if err != nil || t.RootMaterialTypeId == id {
		util.AbortNewResultErrorOfClient(c,
			errors.New(err.Error()+"此物料类别不能删除"))
		return
	}
	err = dao.MaterialTypeDelete(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, 1))
}

func MaterialTypeUpdate(c *gin.Context) {
	var err error
	var form bean.MaterialType
	MaterialTypeId := c.Param("id")
	id, err := strconv.ParseInt(MaterialTypeId, 10, 64)
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
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	//check code
	//parent
	parentMaterialType, err := dao.MaterialTypeRead(TenantId, form.PId)
	if !strings.HasPrefix(form.Code, parentMaterialType.Code) || len(parentMaterialType.Code)+3 != len(form.Code) {
		util.AbortNewResultErrorOfClient(c, errors.New("编码必须满足3-3-3...的格式,且以父级编码为前缀"))
		return
	}

	r, err := dao.MaterialTypeUpdate(form)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	c.JSON(200, util.NewResultOKofWrite(r, 1))
}

func MaterialTypeRead(c *gin.Context) {
	MaterialTypeId := c.Param("id")
	id, err := strconv.ParseInt(MaterialTypeId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.MaterialTypeRead(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func MaterialTypeList(c *gin.Context) {
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.MaterialTypeList(TenantId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}

func MaterialTypeTree(c *gin.Context) {
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	t, _ := dao.TenantRead(TenantId)

	MaterialType, _ := dao.MaterialTypeRead(TenantId, t.RootMaterialTypeId)

	v, err := dao.MaterialTypeTree(TenantId, MaterialType.Code)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	tree := util.TreeOfMaterialType(v)
	c.JSON(200, util.NewResultOKofRead(tree, 1))
}
