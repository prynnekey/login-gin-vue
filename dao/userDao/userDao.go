package userDao

import (
	"fmt"
	"login-vue/global"
	"login-vue/models"
)

// 将一条用户数据保存到数据库
// @return 影响的行数
func Save(user *models.User) int64 {
	row := global.DB.Create(user).RowsAffected
	return row
}

// 判断用户名是否存在
// @return true: 用户名不存在 false:用户名存在
func IsUsernameNotExist(username string) bool {
	user := models.User{}
	// 从数据库中查询数据
	db := global.DB
	db.Where("username = ?", username).First(&user)

	fmt.Printf("user: %v\n", user)

	return user.ID != 0
}

// 判断手机号是否存在
// @return true: 手机号不存在 false:手机号存在
func IsTelNotExist(tel string) bool {
	user := models.User{}
	// 从数据库中查询数据
	db := global.DB
	db.Where("tel = ?", tel).First(&user)

	fmt.Printf("user: %v\n", user)

	return user.ID != 0
}

// 根据用户名查询一条user数据
// @param username 要查询的用户名
// @return 查询到的结果
func GetByUsername(username string) (models.User, error) {
	user := models.User{}
	// select * from user where username = ?
	err := global.DB.Where("username = ?", username).First(&user).Error
	if user.ID == 0 {
		// 没查到 用户不存在
		return user, err
	}

	return user, nil
}
