package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantId, _ := c.MustGet("TENANT_ID").(int64)
		t, _ := dao.TenantRead(tenantId)

		role, _ := c.MustGet("ROLE_ID").(int64)

		if tenantId == 0 && role == 0 {
			c.Next()
			return
		}
		if role != t.RoleAdmin {
			//判断权限
			method := c.Request.Method
			path := c.FullPath()
			entry, err := dao.FindEntry(method, path)
			if entry.ID == 0 || err != nil {
				c.Abort()
				return
			}

			auth, err := dao.FindAuth(role, entry.ID)
			if auth.ID == 0 || err != nil {
				util.AbortNewResultErrorOfServer(c, errors.New("无权操作"))
				return
			}
		} else {
			c.Next()
		}

		return
	}
}
