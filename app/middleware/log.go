package middleware

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"io/ioutil"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		body, _ := ioutil.ReadAll(c.Request.Body)
		log := bean.Log{
			ID:     0,
			UserId: 0,
			Method: c.Request.Method,
			Uri:    c.Request.RequestURI,
			Data:   string(body),
			Ip:     c.ClientIP(),
			Ua:     c.Request.UserAgent(),
		}
		if claims, ok := c.Get("claims"); ok {
			if cc, ok := claims.(*CustomClaims); ok {
				log.UserId = cc.ID
			}
		}
		if _, err := dao.LogCreate(log); err != nil {
			c.Abort()
			return
		}

	}
}
