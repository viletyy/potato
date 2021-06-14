/*
 * @Date: 2021-06-12 21:55:50
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 21:00:47
 * @FilePath: /potato/internal/middleware/access_log.go
 */
package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"go.uber.org/zap"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		global.GO_LOG.With(
			zap.String("request", c.Request.PostForm.Encode()),
			zap.String("response", bodyWriter.body.String()),
			zap.String("trace_id", c.GetString("X-Trace-ID")),
			zap.String("span_id", c.GetString("X-Span-ID")),
		).Sugar().Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
