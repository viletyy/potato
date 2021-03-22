/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-23 00:50:54
 * @FilePath: /potato/routers/router.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	v1 "github.com/viletyy/potato/controller/api/v1"
	_ "github.com/viletyy/potato/docs"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/middleware"
)

var (
	Engine        = gin.Default()
	V1RouterGroup = Engine.Group("../api/v1")
)

func InitRouter() *gin.Engine {
	Engine.Use(gin.Logger())

	Engine.Use(gin.Recovery())

	gin.SetMode(global.GO_CONFIG.App.RunMode)

	Engine.Use(middleware.CORS())

	Engine.GET("/api/v1/auth", v1.GetUserAuth)
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	V1InitModule()

	return Engine
}

func V1InitModule() {
	V1RouterGroup.Use(middleware.JWT())
	users := V1RouterGroup.Group("users")
	users.POST("", v1.AddUser)
	V1InitBasicRouter()
}
