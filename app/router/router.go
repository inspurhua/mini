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
		_api.GET("/menu", api.Entries)
		_auth := _api.Group("")
		_auth.Use(middleware.Auth())
		{
			_auth.GET("/role", api.RoleList)
			_auth.POST("/role", api.RoleCreate)
			_auth.DELETE("/role/:id", api.RoleDelete)
			_auth.PUT("/role/:id", api.RoleUpdate)
			_auth.GET("/role/:id", api.RoleRead)

			_auth.GET("/entry", api.EntryList)
			_auth.POST("/entry", api.EntryCreate)
			_auth.DELETE("/entry/:id", api.EntryDelete)
			_auth.PUT("/entry/:id", api.EntryUpdate)
			_auth.GET("/entry/:id", api.EntryRead)

		}
	}

	return r
}
