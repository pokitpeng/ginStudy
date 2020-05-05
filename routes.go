package main

import (
	"ginStudy/controller"
	"ginStudy/middleware"

	"github.com/gin-gonic/gin"
)

// 路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) //使用中间件保护用户信息接口
	return r
}
