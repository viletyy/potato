/*
 * @Date: 2021-06-10 18:54:19
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 15:41:34
 * @FilePath: /potato/internal/dao/auth.go
 */
package dao

import "github.com/viletyy/potato/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}

	return auth.Get(d.Engine)
}
