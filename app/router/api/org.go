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

func OrgCreate(c *gin.Context) {
	var err error
	var form bean.Org

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
	parentOrg, err := dao.OrgRead(TenantId, form.PId)
	if !strings.HasPrefix(form.Code, parentOrg.Code) || len(parentOrg.Code)+3 != len(form.Code) {
		util.AbortNewResultErrorOfClient(c, errors.New("编码必须满足3-3-3...的格式,且以父级编码为前缀"))
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

	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	t, err := dao.TenantRead(TenantId)
	if err != nil || t.RootOrgId == id {
		util.AbortNewResultErrorOfClient(c,
			errors.New(err.Error()+"此组织不能删除"))
		return
	}
	err = dao.OrgDelete(TenantId, id)
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
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	form.TenantId = TenantId
	//check code
	//parent
	parentOrg, err := dao.OrgRead(TenantId, form.PId)
	if !strings.HasPrefix(form.Code, parentOrg.Code) || len(parentOrg.Code)+3 != len(form.Code) {
		util.AbortNewResultErrorOfClient(c, errors.New("编码必须满足3-3-3...的格式,且以父级编码为前缀"))
		return
	}

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
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.OrgRead(TenantId, id)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, 1))
}

func OrgList(c *gin.Context) {
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	r, err := dao.OrgList(TenantId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}

func OrgTree(c *gin.Context) {
	//roleId, _ := c.MustGet("ROLE_ID").(int64)
	orgId, _ := c.MustGet("ORG_ID").(int64)
	TenantId, _ := c.MustGet("TENANT_ID").(int64)
	org, _ := dao.OrgRead(TenantId, orgId)

	v, err := dao.OrgTree(TenantId, org.Code)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	tree := util.TreeOfOrg(v)
	c.JSON(200, util.NewResultOKofRead(tree, 1))
}
