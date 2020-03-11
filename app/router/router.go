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
		_api.POST("/refresh", api.Refresh)

	}
	_api.Use(middleware.JWTAuth(), middleware.Auth())
	{
		_api.POST("/file", api.Refresh)
	}

	return r
}
