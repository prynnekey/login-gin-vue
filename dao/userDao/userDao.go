package userDao

import (
	"fmt"
	"login-vue/models"

	"gorm.io/gorm"
)

// 判断用户名是否存在
// @return true: 用户名不存在 false:用户名存在
func IsUsernameNotExist(db *gorm.DB, username string) bool {
	user := models.User{}
	// 从数据库中查询数据
	db.Where("username = ?", username).First(&user)

	fmt.Printf("user: %v\n", user)

	return user.ID != 0
}

// 判断手机号是否存在
// @return true: 手机号不存在 false:手机号存在
func IsTelNotExist(db *gorm.DB, tel string) bool {
	user := models.User{}
	// 从数据库中查询数据
	db.Where("tel = ?", tel).First(&user)

	fmt.Printf("user: %v\n", user)

	return user.ID != 0
}
