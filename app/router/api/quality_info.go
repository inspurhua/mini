package api

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
)

func QualityInfoList(c *gin.Context) {
	keyId := c.Param("id")
	id, err := strconv.ParseInt(keyId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	tenantId, _ := c.MustGet("TENANT_ID").(int64)

	r, err := dao.QualityInfoList(tenantId, id)

	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofRead(r, len(r)))
	return
}
func QualityInfoUpdate(c *gin.Context) {
	keyId := c.Param("id")
	id, err := strconv.ParseInt(keyId, 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	tenantId, _ := c.MustGet("TENANT_ID").(int64)

	var data []bean.QualityInfo

	err = c.Bind(&data)

	if err != nil {
		util.AbortNewResultErrorOfClient(c, err)
		return
	}

	err = dao.QualityInfoUpdate(tenantId, id, data)

	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite("", 1))
	return
}
