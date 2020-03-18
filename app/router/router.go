package router

import (
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/config"
	"huage.tech/mini/app/middleware"
	"huage.tech/mini/app/router/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.GET("/test", api.Test)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(config.RunMode)

	_api := r.Group("/api")
	_api.POST("/login", api.Login)
	_api.Use(middleware.JWTToken())
	{
		//刷新token
		_api.POST("/change", api.ChangPassword)
		_api.POST("/file", api.FileCreate)
		_api.GET("/file/:id", api.FileRead)
		_api.GET("/refresh", api.Refresh)
		_api.GET("/menu", api.Entries)
		_api.GET("/orgtree", api.OrgTree)

		_auth := _api.Group("")
		_auth.Use(middleware.Auth(), middleware.Logger())
		{
			_auth.GET("/role", api.RoleList)
			_auth.POST("/role", api.RoleCreate)
			_auth.DELETE("/role/:id", api.RoleDelete)
			_auth.PUT("/role/:id", api.RoleUpdate)
			_auth.GET("/role/:id", api.RoleRead)

			_auth.GET("/role/:id/auth", api.AuthList)
			_auth.PUT("/role/:id/auth", api.AuthUpdate)

			_auth.GET("/entry", api.EntryList)
			_auth.POST("/entry", api.EntryCreate)
			_auth.DELETE("/entry/:id", api.EntryDelete)
			_auth.PUT("/entry/:id", api.EntryUpdate)
			_auth.GET("/entry/:id", api.EntryRead)

			_auth.GET("/org", api.OrgList)
			_auth.POST("/org", api.OrgCreate)
			_auth.DELETE("/org/:id", api.OrgDelete)
			_auth.PUT("/org/:id", api.OrgUpdate)
			_auth.GET("/org/:id", api.OrgRead)

			_auth.GET("/user", api.UserList)
			_auth.POST("/user", api.UserCreate)
			_auth.DELETE("/user/:id", api.UserDelete)
			_auth.PUT("/user/:id", api.UserUpdate)
			_auth.GET("/user/:id", api.UserRead)

			_auth.GET("/log", api.LogList)
		}
	}

	return r
}
