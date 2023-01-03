package router

import (
	"login-vue/controller/userController"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		// 用户注册
		auth.POST("/register", userController.Register)
		// 用户登录
		auth.POST("/login", userController.Login)
	}

	r.Run()
}
