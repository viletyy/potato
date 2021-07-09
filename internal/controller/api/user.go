/*
 * @Date: 2021-07-09 14:44:50
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:45:52
 * @FilePath: /potato/internal/controller/api/user.go
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode"
)

// @Summary 用户注册
// @Description
// @Accept mpfd
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param nickname formData string false "昵称"
// @Success 200 {object} errcode.Error "请求成功"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	param := service.UserRegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	dbUser, err := svc.RegisterUser(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.RegisterUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorRegisterUserError)
		return
	}

	response.ToResponse(dbUser)
}

// @Summary 用户登录
// @Description
// @Accept mpfd
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} errcode.Error "请求成功"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	param := service.UserLoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	dbUser, err := svc.LoginUser(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.LoginUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorLoginUserError)
		return
	}

	token, err := app.GenerateToken(dbUser.Username, dbUser.Password)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"username": dbUser.Username,
		"is_admin": dbUser.IsAdmin,
		"token":    token,
	})
}
