package models

import (
	"crypto/md5"
	"fmt"
	"github.com/viletyy/potato/pkg/util"
	"io"
	"log"
)

type User struct {
	util.Model

	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func CheckUser(username, password string) bool {
	var user User
	util.DB.Select("id").Where(User{Username: username, Password: GetSecretPassword(password)}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func GetSecretPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	final := fmt.Sprintf("%x", h.Sum(nil))
	return final
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (users []User) {
	util.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)
	return
}

func GetUsersTotal(maps interface{}) (count int) {
	util.DB.Model(&User{}).Where(maps).Count(&count)
	return
}

func ExistUserByUsername(username string) bool {
	var user User
	util.DB.Select("id").Where("name = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func ExistUserById(id int) bool {
	var user User
	util.DB.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func AddUser(data map[string]interface{}) bool {
	user := &User{}
	error := util.FillStruct(data, user)
	if error != nil {
		log.Printf("Fill Struct is Fail")
	}
	util.DB.Create(user)

	return true
}

func EditUser(id int, data interface{}) bool {
	util.DB.Model(&User{}).Where("id = ?", id).Update(data)

	return true
}

func DeleteUser(id int) bool {
	util.DB.Where("id = ?", id).Delete(&User{})

	return true
}