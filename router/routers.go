package router

import (
	"login-vue/controller/userController"
	"login-vue/middleware/authMiddleware"
	"login-vue/middleware/corsMiddleware"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()

	// 开启全局中间件 处理跨域请求
	r.Use(corsMiddleware.CORSMiddleware())

	auth := r.Group("/auth")
	{
		// 用户注册
		auth.POST("/register", userController.Register)
		// 用户登录
		auth.POST("/login", userController.Login)

		// 只有授权才可以访问
		auth.GET("/info", authMiddleware.AuthMiddleware(), userController.Info)
	}

	r.Run()
}
