/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:20:33
 * @FilePath: /potato/internal/model/user.go
 */
package model

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"-"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"is_admin" gorm:"default: false"`
}
