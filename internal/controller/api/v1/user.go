/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:32:27
 * @FilePath: /potato/internal/controller/api/v1/user.go
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode"
	"github.com/viletyy/yolk/convert"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @Summary 用户列表
// @Tags users
// @Description
// @Accept json
// @Produce json
// @Param token header string true "auth token"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param username query string false "用户名" maxlength(30)
// @Param nickname query string false "昵称" maxlength(30)
// @Success 200 {object} model.User "请求成功"
// @Router /v1/users [get]
func (user User) List(c *gin.Context) {
	param := service.UserListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	paper := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	total, err := svc.CountUser(&service.CountUserRequest{Username: param.Username, Nickname: param.Nickname})
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.CountUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountUserFail)
		return
	}

	users, err := svc.GetUserList(&param, &paper)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.GetUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}

	response.ToResponseList(users, total)
}

// @Summary 用户
// @Tags users
// @Description
// @Accept json
// @Produce json
// @Param token header string true "auth token"
// @Param id path int true "用户 ID"
// @Success 200 {object} model.User "请求成功"
// @Router /v1/users/{id} [get]
func (user User) Get(c *gin.Context) {
	userID, err := convert.StrTo(c.Param("id")).Int64()
	response := app.NewResponse(c)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	param := service.UserRequest{ID: userID}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	dbUser, err := svc.GetUser(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.GetUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserFail)
		return
	}

	response.ToResponse(dbUser)
}

// @Summary 新增用户
// @Tags users
// @Description
// @Accept mpfd
// @Produce json
// @Param token header string true "auth token"
// @Param username formData string true "用户名" minlength(1) maxlength(30)
// @Param password formData string true "密码" minlength(1) maxlength(18)
// @param nickname formData string false "昵称" maxlength(30)
// @Success 200 {object} model.User "请求成功"
// @Router /v1/users [post]
func (user User) Create(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	dbUser, err := svc.CreateUser(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(dbUser)
}

// @Summary 更新用户
// @Tags users
// @Description
// @Accept mpfd
// @Produce json
// @Param token header string true "auth token"
// @Param id path int true "用户ID"
// @Param username formData string false "用户名" maxlength(30)
// @Param password formData string false "密码" maxlength(18)
// @param nickname formData string false "昵称" maxlength(30)
// @Success 200 {object} model.User "请求成功"
// @Router /v1/users/{id} [patch]
func (user User) Update(c *gin.Context) {
	userID, err := convert.StrTo(c.Param("id")).Int64()
	response := app.NewResponse(c)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	param := service.UpdateUserRequest{ID: userID}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs)
		return
	}
	svc := service.New(c.Request.Context())
	dbUser, err := svc.UpdateUser(&param)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GO_LOG.Sugar().Errorf("svc.UpdateUser err: %v", err)
			response.ToErrorResponse(errcode.ErrorGetUserFail)
			return
		} else {
			global.GO_LOG.Sugar().Errorf("svc.UpdateUser err: %v", err)
			response.ToErrorResponse(errcode.ErrorUpdateUserFail)
			return
		}
	}

	response.ToResponse(dbUser)
}

// @Summary 删除用户
// @Tags users
// @Description
// @Accept json
// @Produce json
// @Param token header string true "auth token"
// @Param id path int true "用户ID"
// @Success 200 {object} model.User "请求成功"
// @Router /v1/users/{id} [delete]
func (user User) Delete(c *gin.Context) {
	userID, err := convert.StrTo(c.Param("id")).Int64()
	response := app.NewResponse(c)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	param := service.DeleteUserRequest{ID: userID}
	svc := service.New(c.Request.Context())
	dbUser, err := svc.DeleteUser(&param)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GO_LOG.Sugar().Errorf("svc.DeleteUser err: %v", err)
			response.ToErrorResponse(errcode.ErrorGetUserFail)
			return
		} else {
			global.GO_LOG.Sugar().Errorf("svc.DeleteUser err: %v", err)
			response.ToErrorResponse(errcode.ErrorDeleteUserFail)
			return
		}
	}

	errcode.Success.Data = dbUser
	response.ToErrorResponse(errcode.Success)
}
