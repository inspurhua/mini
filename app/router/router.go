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
	r.GET("/d", api.DeviceList)
	r.POST("/d/:device_id", api.DeviceCommand)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(config.RunMode)

	_api := r.Group("/api")
	_api.POST("/login", middleware.Logger(), api.Login)
	_api.Use(middleware.JWTToken())
	{
		//刷新token
		_api.PUT("/change", api.ChangPassword)
		_api.POST("/file", api.FileCreate)
		_api.GET("/file/:id", api.FileRead)
		_api.GET("/refresh", api.Refresh)
		_api.GET("/menu", api.Entries)
		_api.GET("/orgtree", api.OrgTree)

		_auth := _api.Group("")
		_auth.Use(middleware.Auth())
		{
			_auth.GET("/log", api.LogList)
		}
		_auth.Use(middleware.Auth(), middleware.Logger())
		{
			_auth.GET("/tenant", api.TenantList)
			_auth.POST("/tenant", api.TenantCreate)
			_auth.PUT("/tenant/:id", api.TenantUpdate)
			_auth.GET("/tenant/:id", api.TenantRead)

			_auth.GET("/role", api.RoleList)
			_auth.POST("/role", api.RoleCreate)
			_auth.DELETE("/role/:id", api.RoleDelete)
			_auth.PUT("/role/:id", api.RoleUpdate)
			_auth.GET("/role/:id", api.RoleRead)

			_auth.GET("/role/:id/auth", api.AuthList)
			_auth.PUT("/role/:id/auth", api.AuthUpdate)

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

			_auth.GET("/material_type", api.MaterialTypeList)
			_auth.POST("/material_type", api.MaterialTypeCreate)
			_auth.DELETE("/material_type/:id", api.MaterialTypeDelete)
			_auth.PUT("/material_type/:id", api.MaterialTypeUpdate)
			_auth.GET("/material_type/:id", api.MaterialTypeRead)

			_auth.GET("/material", api.MaterialList)
			_auth.POST("/material", api.MaterialCreate)
			_auth.DELETE("/material/:id", api.MaterialDelete)
			_auth.PUT("/material/:id", api.MaterialUpdate)
			_auth.GET("/material/:id", api.MaterialRead)

			_auth.GET("/quality_info/:id", api.QualityInfoList)
			_auth.POST("/quality_info/:id", api.QualityInfoUpdate)
		}
	}

	return r
}
