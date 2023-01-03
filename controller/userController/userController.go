package userController

import (
	"log"
	"login-vue/common/code"
	"login-vue/common/response"
	"login-vue/dao/userDao"
	"login-vue/dto/userDto"
	"login-vue/models"
	"login-vue/utils"
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
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "手机号格式不正确")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "密码不能少于6位")
		return
	}

	if len(username) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "用户名为空")
		return
	}

	// 判断两次密码是否一致
	if password != confirmPassword {
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "两次密码不一致")
		return
	}

	// 判断用户名是否存在
	if userDao.IsUsernameNotExist(username) {
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "用户名已存在")
		return
	}

	// 判断手机号是否存在
	if userDao.IsTelNotExist(tel) {
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "手机号已存在")
		return
	}

	// 注册成功 将数据保存到数据库
	user := &models.User{
		Username: username,
		Password: password,
		Tel:      tel,
	}

	// 插入数据
	i := userDao.Save(user)
	if i == 0 {
		response.Response(ctx, http.StatusInternalServerError, code.SERVER_ERROR, nil, "服务器故障啦！注册失败")
	}

	// 返回信息
	response.Response(ctx, http.StatusOK, code.POST_OK, user, "注册成功!")
}

// 用户登录
func Login(ctx *gin.Context) {
	// 1. 获取用户名和密码
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 2. 根据用户名查询数据
	user, err := userDao.GetByUsername(username)
	if err != nil {
		// 3. 没查到 用户不存在
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "用户名不存在")
		return
	}

	// 4. 查到 对比密码是否正确
	if password != user.Password {
		// 5. 密码不正确 返回错误信息
		response.Response(ctx, http.StatusUnprocessableEntity, code.POST_ERROR, nil, "密码错误")
		return
	}
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	// 	// 5. 密码不正确 返回错误信息
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
	// 	return
	// }
	// 6. 密码正确 发放token
	token, err := utils.GetToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, code.SERVER_ERROR, nil, "系统异常,登录失败")
		log.Printf("token generate error : %v", err)
		return
	}

	// 7. 返回登录成功
	response.Success(ctx, code.POST_OK, gin.H{"token": token}, "登录成功")
}

// 用户登录成功后的信息
func Info(ctx *gin.Context) {
	// 从上下文中获取user
	_user, _ := ctx.Get("user")
	_u := _user.(models.User)
	userDto := userDto.NewUserDto(&_u)
	response.Success(ctx, code.GET_OK, gin.H{"user": userDto}, "")
}
