package authMiddleware

import (
	"login-vue/dao/userDao"
	"login-vue/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 授权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从authorization header中获取token
		tokenString := ctx.GetHeader("Authorization")

		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			// 验证不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 后续不再进行
			ctx.Abort()
			return
		}

		// 验证通过
		// 解析token
		tokenString = tokenString[7:]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			// 解析不通过
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 解析通过 从Claim中获取userId
		userId := claims.UserID
		// 根据userId查询user
		user, err := userDao.GetById(userId)
		if err != nil {
			// 用户不存在
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在 将用户写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
