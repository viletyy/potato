/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 19:02:31
 * @FilePath: /potato/internal/routers/router.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/viletyy/potato/docs"
	"github.com/viletyy/potato/global"
	v1 "github.com/viletyy/potato/internal/controller/api/v1"
	"github.com/viletyy/potato/internal/middleware"
)

var (
	Engine        = gin.Default()
	V1RouterGroup = Engine.Group("../api/v1")
)

func InitRouter() *gin.Engine {
	Engine.Use(gin.Logger())

	Engine.Use(gin.Recovery())
	Engine.Use(middleware.Translations())

	gin.SetMode(global.GO_CONFIG.App.RunMode)

	Engine.Use(middleware.CORS())

	Engine.POST("/api/v1/auth", v1.GetAuth)
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	V1InitModule()

	return Engine
}

func V1InitModule() {
	V1RouterGroup.Use(middleware.JWT())
	V1InitBasicRouter()
}
