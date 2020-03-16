package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/util"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if roleId, ok := c.Get("ROLE_ID"); ok {
			if role, ok := roleId.(int64); ok {
				if role != 1 {
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
				}
			}
			c.Next()
		}
		c.Abort()
		return
	}
}
