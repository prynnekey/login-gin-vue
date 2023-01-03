package config

import (
	"login-vue/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库
func InitConfig() {
	dsn := "root:prynnekey@tcp(127.0.0.1:3306)/gin_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 查看sql语句日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	// 自动创建user表
	// db.AutoMigrate(&User{})

	global.DB = db
}
