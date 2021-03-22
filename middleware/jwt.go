/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-23 00:50:00
 * @FilePath: /potato/middleware/jwt.go
 */
package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "请求参数错误",
			})
			c.Abort()
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "token验证失败",
				})
				c.Abort()
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "token已超时",
				})
				c.Abort()
			}
			if claims != nil {
				userId := claims.UserId
				loginUUID := claims.StandardClaims.Id
				val, _ := global.GO_REDIS.Get("login:" + loginUUID).Result()
				if val != strconv.Itoa(int(userId)) {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "token鉴权失败",
					})
					c.Abort()
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "token鉴权失败",
				})
				c.Abort()
			}
		}

		c.Next()
	}
}
