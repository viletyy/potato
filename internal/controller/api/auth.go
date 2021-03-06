/*
 * @Date: 2021-06-10 18:58:25
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 21:02:33
 * @FilePath: /potato/internal/controller/api/auth.go
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode"
)

// @Summary 鉴权验证
// @Description
// @Accept mpfd
// @Produce json
// @Param app_key formData string true "app key"
// @Param app_secret formData string true "app secret"
// @Success 200 {object} errcode.Error "请求成功"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	auth, err := svc.CheckAuth(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"app_key":    auth.AppKey,
		"app_secret": auth.AppSecret,
		"token":      token,
	})
}
