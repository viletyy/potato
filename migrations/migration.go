package migrations

import (
	"fmt"
	"log"
	"os"

	"github.com/viletyy/potato/global"
	"gorm.io/gorm"
)

const minDBVersion = 0

type Migration interface {
	Description() string
	Migrate(*gorm.DB) error
}

type migration struct {
	description string
	migrate     func(*gorm.DB) error
}

func NewMigration(desc string, fn func(*gorm.DB) error) Migration {
	return &migration{desc, fn}
}

func (m *migration) Description() string {
	return m.description
}

func (m *migration) Migrate(db *gorm.DB) error {
	return m.migrate(db)
}

type Version struct {
	ID      int64 `gorm:"primary_key"`
	Version int64
}

var migrations = []Migration{
	NewMigration("create table users", createTableUsers),
	NewMigration("create table auths", createTableAuths),
	NewMigration("create table vendors", createTableVendors),
}

func GetCurrentDBVersion(db *gorm.DB) (int64, error) {
	if err := db.Debug().AutoMigrate(&Version{}); err != nil {
		return -1, fmt.Errorf("db.AutoMigrate: %v", err)
	}

	currentVersion := &Version{ID: 1}
	if err := db.Debug().First(currentVersion).Error; err != nil {
		return -1, fmt.Errorf("db.First: %v", err)
	}

	return currentVersion.Version, nil
}

func ExpectedVersion() int64 {
	return int64(minDBVersion + len(migrations))
}

func EnsureUpTodate(db *gorm.DB) error {
	currentDB, err := GetCurrentDBVersion(db)
	if err != nil {
		return err
	}

	if currentDB < 0 {
		return fmt.Errorf("Database has not been initialised")
	}

	if minDBVersion > currentDB {
		return fmt.Errorf("DB version %d (<= %d) is too old for auto-migration.", currentDB, minDBVersion)
	}

	expected := ExpectedVersion()

	if currentDB != expected {
		return fmt.Errorf(`Current database version %d is not equal to the expected version %d. `, currentDB, expected)
	}

	return nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Version{}); err != nil {
		return fmt.Errorf("db.AutoMigrate: %v", err)
	}

	currentVersion := &Version{ID: 1}
	if err := db.First(currentVersion).Error; err != nil {
		currentVersion.Version = 0
		if err := db.Debug().Create(currentVersion).Error; err != nil {
			return fmt.Errorf("db.Create: %v", err)
		}
	}

	v := currentVersion.Version
	if minDBVersion > v {
		global.GO_LOG.Fatal("Please upgrade the latest code.")
	}

	if int(v-minDBVersion) > len(migrations) {
		msg := fmt.Sprintf("Downgrading database version from '%d' to '%d' is not supported and may result in loss of data integrity.\nIf you really know what you're doing, execute `UPDATE version SET version=%d WHERE id=1;`\n",
			v, minDBVersion+len(migrations), minDBVersion+len(migrations))
		fmt.Fprint(os.Stderr, msg)
		log.Fatal(msg)
		return nil
	}

	for i, m := range migrations[v-minDBVersion:] {
		global.GO_LOG.Sugar().Infof("Migration[%d]: %s", v+int64(i), m.Description())
		if err := m.Migrate(db); err != nil {
			return fmt.Errorf("Migrate: %v", err)
		}
		currentVersion.Version = v + int64(i) + 1
		if err := db.Save(currentVersion).Error; err != nil {
			return err
		}
	}

	return nil
}
