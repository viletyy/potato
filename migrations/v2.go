package migrations

import (
	"fmt"

	"github.com/viletyy/potato/internal/model"
	"gorm.io/gorm"
)

func createTableAuths(db *gorm.DB) (err error) {
	type Auth struct {
		*model.Model

		AppKey    string `json:"app_key"`
		AppSecret string `json:"app_secret"`
	}

	if err := db.Debug().AutoMigrate(&Auth{}); err != nil {
		return fmt.Errorf("migrations: create table auths err: %v", err)
	}

	return nil
}
