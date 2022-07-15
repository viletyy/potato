package migrations

import (
	"fmt"

	"github.com/viletyy/potato/internal/model"
	"gorm.io/gorm"
)

func createTableVendors(db *gorm.DB) (err error) {
	type Vendor struct {
		*model.Model

		Name string `json:"name"`
		Uuid int    `json:"uuid"`
	}

	if err := db.Debug().AutoMigrate(&Vendor{}); err != nil {
		return fmt.Errorf("migrations: create table auths err: %v", err)
	}

	return nil
}
