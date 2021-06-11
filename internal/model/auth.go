/*
 * @Date: 2021-06-10 18:21:37
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 15:41:04
 * @FilePath: /potato/internal/model/auth.go
 */
package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) Get(db *gorm.DB) (auth Auth, err error) {
	if notFound := db.Where("app_key = ? AND app_secret = ?", a.AppKey, a.AppSecret).First(&auth).RecordNotFound(); notFound {
		return a, gorm.ErrRecordNotFound
	}

	return auth, nil
}
