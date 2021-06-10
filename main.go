/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 10:28:27
 * @FilePath: /potato/main.go
 */
package main

import (
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/initialize"
)

// @title Potato Api
// @version 1.0
// @description This is a potato use golang
// @BasePath /api

func main() {
	global.GO_VP = initialize.Viper()
	global.GO_LOG = initialize.Zap()
	global.GO_DB = initialize.Gorm()
	global.GO_REDIS = initialize.Redis()

	defer global.GO_DB.Close()
	defer global.GO_REDIS.Close()

	initialize.RunServer()

}
