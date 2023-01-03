package main

import (
	"login-vue/config"
	"login-vue/router"
)

func main() {
	config.InitConfig()

	router.InitRouters()
}
