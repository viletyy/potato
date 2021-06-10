/*
 * @Date: 2021-06-10 18:21:37
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 21:59:51
 * @FilePath: /potato/internal/model/auth.go
 */
package model

import "github.com/jinzhu/gorm"

type Auth struct {
	Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ?", a.AppKey, a.AppSecret)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
