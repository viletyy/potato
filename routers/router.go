package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/viletyy/potato/controller/api"
	_ "github.com/viletyy/potato/docs"
	"github.com/viletyy/potato/middleware/cors"
	"github.com/viletyy/potato/pkg/setting"
)

var (
	Engine = gin.Default()
	V1RouterGroup = Engine.Group("../api/v1")
)

func InitRouter() *gin.Engine {
	Engine.Use(gin.Logger())

	Engine.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	Engine.Use(cors.CORSMiddleware())

	Engine.GET("/auth", api.GetAuth)
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	V1InitModule()

	return Engine
}

func V1InitModule() {
	V1InitBasicRouter()
}
