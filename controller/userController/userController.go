package userController

import (
	"login-vue/dao/userDao"
	"login-vue/global"
	"login-vue/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	confirmPassword := ctx.PostForm("confirmPassword")
	tel := ctx.PostForm("tel")

	// 参数校验
	if len(tel) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号格式不正确",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	if len(username) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名为空",
		})
		return
	}

	// 判断两次密码是否一致
	if password != confirmPassword {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "两次密码不一致",
		})
		return
	}

	db := global.DB

	// 判断用户名是否存在
	if userDao.IsUsernameNotExist(db, username) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名已存在",
		})
		return
	}

	// 判断手机号是否存在
	if userDao.IsTelNotExist(db, tel) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号已存在",
		})
		return
	}

	// 注册成功 将数据保存到数据库
	user := &models.User{
		Username: username,
		Password: password,
		Tel:      tel,
	}

	// 插入数据
	db.Create(user)

	// 返回信息
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": user,
	})
}

// 用户登录
func Login(ctx *gin.Context) {

}
