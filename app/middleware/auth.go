package middleware

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if claims, ok := c.Get("claims"); ok {
			if cc, ok := claims.(*CustomClaims); ok {
				if cc.Role != 1 {
					//判断权限
					method := c.Request.Method
					path := c.FullPath()
					entry, err := dao.FindEntry(method, path)
					if entry.ID == 0 || err != nil {
						c.Abort()
						return
					}
					auth, err := dao.FindAuth(cc.Role, entry.ID)
					if auth.ID == 0 || err != nil {
						c.Abort()
						return
					}
					c.Next()
				} else {
					c.Next()
				}
			}
		}
		c.Abort()
		return
	}
}
