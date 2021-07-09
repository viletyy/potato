/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:44:37
 * @FilePath: /potato/internal/routers/router.go
 */
package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/viletyy/potato/docs"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/controller/api"
	v1 "github.com/viletyy/potato/internal/controller/api/v1"
	"github.com/viletyy/potato/internal/middleware"
	"github.com/viletyy/potato/pkg/limiter"
)

var (
	Engine         = gin.Default()
	methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.GO_CONFIG.App.RunMode)                                                        // 设置运行模式
	Engine.Use(middleware.AppInfo())                                                                 // 设置app信息
	Engine.Use(middleware.RateLimiter(methodLimiters))                                               // 设置限流控制
	Engine.Use(middleware.ContextTimeout(time.Duration(global.GO_CONFIG.App.DefaultContextTimeout))) // 设置统一超时控制
	Engine.Use(middleware.Translations())                                                            // 设置自定义验证
	Engine.Use(middleware.CORS())                                                                    // 设置跨域
	Engine.Use(middleware.Tracing())
	if global.GO_CONFIG.App.RunMode == "debug" {
		Engine.Use(gin.Logger())   // 设置log
		Engine.Use(gin.Recovery()) // 设置recovery
	} else {
		Engine.Use(middleware.AccessLog())
		Engine.Use(middleware.Recovery())
	}

	Engine.StaticFS("/static", http.Dir(global.GO_CONFIG.App.UploadSavePath))

	Engine.POST("/api/auth", api.GetAuth)
	Engine.POST("/api/user/register", api.UserRegister)
	Engine.POST("/api/user/login", api.UserLogin)
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1RouterGroup := Engine.Group("../api/v1")
	v1RouterGroup.Use(middleware.JWT())

	upload := v1.NewUpload()
	v1RouterGroup.POST("/upload", upload.Create)

	users := v1RouterGroup.Group("/users")
	user := v1.NewUser()
	{
		users.GET("", user.List)
		users.POST("", user.Create)
		users.GET("/:id", user.Get)
		users.PATCH("/:id", user.Update)
		users.DELETE("/:id", user.Delete)
	}
	V1InitBasicRouter()

	return Engine
}
