package router

import (
	"login-vue/controller"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		// 用户注册
		auth.POST("/register", controller.UserRegister)
	}

	r.Run()
}
