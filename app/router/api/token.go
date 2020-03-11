package api

import (
	"errors"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/dao"
	"huage.tech/mini/app/middleware"
	"huage.tech/mini/app/util"
	"time"
)

type LoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,gte=6"`
}

func Login(c *gin.Context) {
	var err error
	var req LoginRequest

	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, util.NewResultErrorOfClient(err))
		c.Abort()
		return
	}

	u := dao.Login(req.Account, req.Password)
	if u.ID == 0 {
		c.JSON(200, util.NewResultErrorOfClient(errors.New("账号或者密码错误,请重试")))
		c.Abort()
		return
	}

	jwt := middleware.NewJWT()
	expire := time.Now().Add(2 * time.Hour).Unix()
	claims := middleware.CustomClaims{
		ID:             u.ID,
		Org:            u.OrgId,
		Role:           u.RoleId,
		StandardClaims: jwt2.StandardClaims{ExpiresAt: expire},
	}

	token, err := jwt.CreateToken(claims)

	c.JSON(200, util.NewResultOKofRead(gin.H{"AccessToken": token, "Expire": expire}, 1))
}

func Refresh(c *gin.Context) {
	jwt := middleware.NewJWT()
	if token, ok := c.Get("token"); ok {
		new, err := jwt.RefreshToken(token.(string))
		if err != nil {
			c.JSON(200, util.NewResultErrorOfClient(errors.New("无法刷新得到新的token")))
			c.Abort()
			return
		}
		expire := time.Now().Add(2 * time.Hour).Unix()
		c.JSON(200, util.NewResultOKofRead(gin.H{"AccessToken": new, "Expire": expire}, 1))
	}
	c.JSON(200, util.NewResultErrorOfClient(errors.New("未提供token")))
	c.Abort()
	return

}
