/*
 * @Date: 2021-06-10 18:55:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:58:09
 * @FilePath: /potato/internal/service/auth.go
 */
package service

import "errors"

type AuthRequest struct {
	AppKey    string `json:"app_key" validate:"required"`
	AppSecret string `json:"app_secret" validate:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}
