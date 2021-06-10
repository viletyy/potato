/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 15:20:29
 * @FilePath: /potato/internal/middleware/jwt.go
 */
package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/pkg"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "请求参数错误",
			})
			c.Abort()
			return
		} else {
			claims, err := pkg.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "token验证失败",
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "token已超时",
				})
				c.Abort()
				return
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
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "token鉴权失败",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
