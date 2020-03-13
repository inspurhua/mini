package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if claims, ok := c.Get("claims"); ok {
			if cc, ok := claims.(*CustomClaims); ok {
				if cc.Role != 1 {
					//判断权限
					//TODO
					method := c.Request.Method
					path := c.FullPath()

					c.Abort()
					return
				} else {
					c.Next()
				}
			}
		}
		c.Abort()
		return
	}
}
