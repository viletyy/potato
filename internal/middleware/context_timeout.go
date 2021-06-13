/*
 * @Date: 2021-06-13 22:35:30
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 22:37:09
 * @FilePath: /potato/internal/middleware/context_timeout.go
 */
package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
