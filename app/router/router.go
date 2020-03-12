package router

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/config"
	"huage.tech/mini/app/middleware"
	"huage.tech/mini/app/router/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(config.RunMode)

	_api := r.Group("/api")
	_api.POST("/login", api.Login)
	_api.Use(middleware.JWTAuth())
	{
		//刷新token
		_api.GET("/refresh", api.Refresh)
		_auth := _api.Group("auth")
		_auth.Use(middleware.Auth())
		{
			_api.GET("/role", api.RoleList)
			_api.POST("/role", api.RoleCreate)
			_api.DELETE("/role/:id", api.RoleDelete)
			_api.PUT("/role/:id", api.RoleUpdate)
			_api.GET("/role/:id", api.RoleRead)
		}
	}

	return r
}
