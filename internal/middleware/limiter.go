/*
 * @Date: 2021-06-13 22:27:24
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 22:29:27
 * @FilePath: /potato/internal/middleware/limiter.go
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode"
	"github.com/viletyy/potato/pkg/limiter"
)

func RateLimiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
