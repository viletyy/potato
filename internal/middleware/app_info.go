/*
 * @Date: 2021-06-13 22:01:30
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 22:02:41
 * @FilePath: /potato/internal/middleware/app_info.go
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
)

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", global.GO_CONFIG.App.Name)
		c.Set("app_version", global.GO_CONFIG.App.Version)
		c.Next()
	}
}
