/*
 * @Date: 2021-07-09 14:29:44
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:30:11
 * @FilePath: /potato/internal/service/user.go
 */

package service

import (
	"github.com/viletyy/potato/internal/model"
	"github.com/viletyy/potato/pkg/app"
)

type UserRegisterRequest struct {
	Username string `form:"username" json:"username" validate:"max=30"`
	Password string `form:"password" json:"password" validate:"max=18"`
	Nickname string `form:"nickname" json:"nickname" validate:"max=30"`
}

type UserLoginRequest struct {
	Username string `form:"username" json:"username" validate:"max=30"`
	Password string `form:"password" json:"password" validate:"max=18"`
}

type UserListRequest struct {
	Username string `form:"username" json:"username" validate:"max=30"`
	Nickname string `form:"nickname" json:"nickname" validate:"max=30"`
}

type CountUserRequest struct {
	Username string `form:"username" json:"username" validate:"max=30"`
	Password string `form:"password" json:"password" validate:"max=18"`
	Nickname string `form:"nickname" json:"nickname" validate:"max=30"`
}

type UserRequest struct {
	ID int64 `form:"id" json:"id" validate:"required,gte=1"`
}

type CreateUserRequest struct {
	Username string `form:"username" json:"username" validate:"required,max=30"`
	Password string `form:"password" json:"password" validate:"required,max=18"`
	Nickname string `form:"nickname" json:"nickname" validate:"max=30"`
}

type UpdateUserRequest struct {
	ID       int64  `form:"id" json:"id" validate:"required,gte=1"`
	Username string `form:"username" json:"username" validate:"max=30"`
	Password string `form:"password" json:"password" validate:"max=18"`
	Nickname string `form:"nickname" json:"nickname" validate:"max=30"`
}

type DeleteUserRequest struct {
	ID int64 `form:"id" json:"id" validate:"required,gte=1"`
}

func (svc *Service) RegisterUser(param *UserRegisterRequest) (model.User, error) {
	return svc.Dao.RegisterUser(param.Username, param.Password, param.Nickname)
}

func (svc *Service) LoginUser(param *UserLoginRequest) (model.User, error) {
	return svc.Dao.LoginUser(param.Username, param.Password)
}

func (svc *Service) CountUser(param *CountUserRequest) (int64, error) {
	return svc.Dao.CountUser(param.Username, param.Nickname)
}

func (svc *Service) GetUserList(param *UserListRequest, pager *app.Pager) ([]model.User, error) {
	return svc.Dao.GetUserList(param.Username, param.Nickname, pager.Page, pager.PageSize)
}

func (svc *Service) GetUser(param *UserRequest) (model.User, error) {
	return svc.Dao.GetUser(param.ID)
}

func (svc *Service) CreateUser(param *CreateUserRequest) (model.User, error) {
	return svc.Dao.CreateUser(param.Username, param.Password, param.Nickname)
}

func (svc *Service) UpdateUser(param *UpdateUserRequest) (model.User, error) {
	return svc.Dao.UpdateUser(param.ID, param.Username, param.Password, param.Nickname)
}

func (svc *Service) DeleteUser(param *DeleteUserRequest) (model.User, error) {
	return svc.Dao.DeleteUser(param.ID)
}
