/*
 * @Date: 2021-06-12 22:29:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-12 23:10:53
 * @FilePath: /potato/internal/middleware/recovery.go
 */
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/email"
	"github.com/viletyy/potato/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.GO_CONFIG.Email.Host,
		Port:     global.GO_CONFIG.Email.Port,
		IsSSL:    global.GO_CONFIG.Email.IsSSL,
		UserName: global.GO_CONFIG.Email.UserName,
		Password: global.GO_CONFIG.Email.Password,
		From:     global.GO_CONFIG.Email.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.GO_LOG.Sugar().Errorf("panic recover err: %v", err)

				err := defaultMailer.SendMail(
					global.GO_CONFIG.Email.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.GO_LOG.Sugar().Panicf("mail.SendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
