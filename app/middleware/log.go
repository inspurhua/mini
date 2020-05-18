package middleware

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/bean"
	"huage.tech/mini/app/dao"
	"io/ioutil"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		body, _ := ioutil.ReadAll(c.Request.Body)
		log := bean.Log{
			ID:       0,
			UserId:   0,
			Method:   c.Request.Method,
			Uri:      c.Request.RequestURI,
			Data:     string(body),
			Ip:       c.ClientIP(),
			Ua:       c.Request.UserAgent(),
			CreateAt: time.Now(),
		}
		if uId, ok := c.Get("UID"); ok {
			if id, ok := uId.(int64); ok {
				log.UserId = id
			}
		}

		tId, _ := c.MustGet("TENANT_ID").(int64)
		log.TenantId = tId

		if _, err := dao.LogCreate(log); err != nil {
			c.Abort()
			return
		}

	}
}
