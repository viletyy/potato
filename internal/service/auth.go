/*
 * @Date: 2021-06-10 18:55:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 15:43:31
 * @FilePath: /potato/internal/service/auth.go
 */
package service

import "github.com/viletyy/potato/internal/model"

type AuthRequest struct {
	AppKey    string `form:"app_key" validate:"required"`
	AppSecret string `form:"app_secret" validate:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) (model.Auth, error) {
	return svc.dao.GetAuth(param.AppKey, param.AppSecret)
}
