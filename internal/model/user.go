/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:31:04
 * @FilePath: /potato/internal/model/user.go
 */
package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model

	Username string `json:"username"`
	Password string `json:"-"`
	Nickname string `json:"nickname"`
	IsAdmin  bool   `json:"is_admin" gorm:"default: false"`
}

func (u User) Count(db *gorm.DB) (int, error) {
	var count int
	if u.Username != "" {
		db = db.Where("username = ?", u.Username)
	}
	if u.Nickname != "" {
		db = db.Where("nickname = ?", u.Nickname)
	}
	if err := db.Model(&u).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) (users []User, err error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if u.Username != "" {
		db = db.Where("name = ?", u.Username)
	}
	if u.Nickname != "" {
		db = db.Where("nickname = ?", u.Nickname)
	}

	if err = db.Find(&users).Error; err != nil {
		return nil, err
	}

	return
}

func (u User) GetByUsernameAndPassword(db *gorm.DB) (user User, err error) {
	if notFound := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).RecordNotFound(); notFound {
		return u, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (u User) Get(db *gorm.DB) (user User, err error) {
	if notFound := db.Where("id = ?", u.ID).First(&user).RecordNotFound(); notFound {
		return u, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	err := db.Model(&User{}).Update(u).Error
	return err
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Where("id = ?", u.ID).Delete(u).Error
}
