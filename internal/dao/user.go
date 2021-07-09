/*
 * @Date: 2021-07-09 14:30:19
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:30:40
 * @FilePath: /potato/internal/dao/user.go
 */
package dao

import (
	"github.com/viletyy/potato/internal/model"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/yolk/crypt"
)

func (d *Dao) RegisterUser(username string, password string, nickname string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: crypt.Md5Encode(password),
		Nickname: nickname,
	}

	return user, user.Create(d.Engine)
}

func (d *Dao) LoginUser(username string, password string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: crypt.Md5Encode(password),
	}

	return user.GetByUsernameAndPassword(d.Engine)
}

func (d *Dao) CountUser(username, nickname string) (int, error) {
	vendor := model.User{Username: username, Nickname: nickname}
	return vendor.Count(d.Engine)
}

func (d *Dao) GetUserList(username, nickname string, page, pageSize int) ([]model.User, error) {
	user := model.User{Username: username, Nickname: nickname}
	pageOffset := app.GetPageOffset(page, pageSize)
	return user.List(d.Engine, pageOffset, pageSize)
}

func (d *Dao) GetUser(id int64) (model.User, error) {
	user := model.User{
		Model: &model.Model{ID: id},
	}

	return user.Get(d.Engine)
}

func (d *Dao) CreateUser(username, password, nickname string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: crypt.Md5Encode(password),
		Nickname: nickname,
	}

	return user, user.Create(d.Engine)
}

func (d *Dao) UpdateUser(id int64, username, password, nickname string) (model.User, error) {

	user := model.User{
		Username: username,
		Nickname: nickname,
		Model:    &model.Model{ID: id},
	}

	dbUser, err := user.Get(d.Engine)

	if password == "" {
		user.Password = dbUser.Password
	} else {
		user.Password = crypt.Md5Encode(password)
	}

	if err != nil {
		return user, err
	}

	return dbUser, user.Update(d.Engine)
}

func (d *Dao) DeleteUser(id int64) (model.User, error) {
	user := model.User{
		Model: &model.Model{ID: id},
	}

	dbUser, err := user.Get(d.Engine)

	if err != nil {
		return user, err
	}

	return dbUser, dbUser.Delete(d.Engine)
}
