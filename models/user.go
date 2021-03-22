/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 23:51:55
 * @FilePath: /potato/models/user.go
 */
package models

import (
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/utils"
	"github.com/viletyy/potato/utils/crypt"
)

type UserSearch struct {
	User
	utils.PageInfo
}

type User struct {
	global.Model

	Username string `json:"username" binding:"required"`
	Password string `json:"-" binding:"required"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"is_admin" gorm:"default: false"`
}

func GetUsers(search *UserSearch) (searchResult utils.SearchResult, err error) {
	var users []User
	offset := search.PageInfo.PageSize * (search.PageInfo.Page - 1)
	limit := search.PageInfo.Page
	db := global.GO_DB.Where(search.User)
	err = db.Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return
	}
	searchResult.Page = search.PageInfo.Page
	searchResult.PageSize = search.PageInfo.PageSize
	searchResult.List = users
	return
}

func GetUserById(id int) (user User, err error) {
	err = global.GO_DB.Where("id = ?", id).First(&user).Error
	return
}

func GetUserByUsername(username string) (user User, err error) {
	err = global.GO_DB.Where("username = ?", username).First(&user).Error
	return
}

func ExistUserById(id int64) bool {
	var user User
	global.GO_DB.Where("id = ?", id).First(&user)

	return user.ID > 0
}

func ExistUserByUsername(username string) bool {
	var user User
	global.GO_DB.Where("username = ?", username).First(&user)

	return user.ID > 0
}

func CreateUser(user User) (err error) {
	err = global.GO_DB.Create(&user).Error

	return
}

func UpdateUser(user *User) (err error) {
	err = global.GO_DB.Save(&user).Error
	return
}

func DeleteUser(user *User) (err error) {
	err = global.GO_DB.Delete(&user).Error
	return
}

func (user *User) CheckPassword(password string) bool {
	return crypt.Md5Check(password, user.Password)
}
