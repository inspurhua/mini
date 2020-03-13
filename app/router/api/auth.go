package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
	"strconv"
	"strings"
)

func AuthList(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	tree := c.DefaultQuery("tree", "0")

	entries, err := dao.AuthList(roleId)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	if tree == "1" {
		entries = util.MenuTree(entries)
	}
	c.JSON(200, util.NewResultOKofRead(entries, len(entries)))
	return
}

func AuthUpdate(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}

	ids, ok := c.GetPostForm("ids")
	if !ok {
		util.AbortNewResultErrorOfClient(c, errors.New("ids参数是以逗号分割的"))
		return
	}
	idSlice := strings.Split(ids, ",")

	var idInt64 []int64
	for _, idStr := range idSlice {
		entryId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			util.AbortNewResultErrorOfClient(c, errors.New("ids参数是以逗号分割的"))
			return
		}
		idInt64 = append(idInt64, entryId)
	}
	err = dao.AuthUpdate(roleId, idInt64)
	if err != nil {
		util.AbortNewResultErrorOfServer(c, err)
		return
	}
	c.JSON(200, util.NewResultOKofWrite(nil, len(idInt64)))
}
