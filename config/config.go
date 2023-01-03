package config

import (
	"fmt"
	"login-vue/global"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库
func InitConfig() {
	// 读取yml文件的数据
	initViper()

	// 初始化数据库
	initDB()
}

func initViper() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(workDir + "/config") // path to look for the config file in
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
}

func initDB() {
	host := viper.GetString("databases.mysql.host")
	port := viper.GetString("databases.mysql.port")
	username := viper.GetString("databases.mysql.username")
	password := viper.GetString("databases.mysql.password")
	dbname := viper.GetString("databases.mysql.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
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
