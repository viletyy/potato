package migrations

import (
	"fmt"

	"github.com/viletyy/potato/internal/model"
	"gorm.io/gorm"
)

func createTableUsers(db *gorm.DB) (err error) {
	type User struct {
		*model.Model

		Username string `json:"username"`
		Password string `json:"-"`
		Nickname string `json:"nickname"`
		IsAdmin  bool   `json:"is_admin" gorm:"default: false"`
	}

	if err := db.Debug().AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("migrations: create table users err: %v", err)
	}

	return nil
}
