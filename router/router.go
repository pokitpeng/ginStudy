package router

import (
	"ginStudy/controller"
	"ginStudy/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 收集路由
func CollectRoute(routes *gin.Engine) *gin.Engine {
	// 路由分组，用户api
	v1 := routes.Group("/api/v1/auth")
	{
		v1.POST("/register", controller.Register)
		v1.POST("/login", controller.Login)
		v1.GET("/info", middleware.AuthMiddleware(), controller.Info) //使用中间件保护用户信息接口
	}

	return routes
}
