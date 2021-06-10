/*
 * @Date: 2021-03-22 10:12:38
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 17:52:06
 * @FilePath: /potato/initialize/gorm.go
 */
package initialize

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/model"
	"github.com/viletyy/potato/internal/model/basic"
)

func Gorm() *gorm.DB {
	switch global.GO_CONFIG.Database.Type {
	case "mysql":
		return GormMysql()
	case "postgresql":
		return GormPostgresql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.GO_CONFIG.Database.User, global.GO_CONFIG.Database.Password, global.GO_CONFIG.Database.Host, global.GO_CONFIG.Database.Port, global.GO_CONFIG.Database.Name))
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Mysql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormPostgresql() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable password=%s", global.GO_CONFIG.Database.Host, global.GO_CONFIG.Database.User, global.GO_CONFIG.Database.Name, global.GO_CONFIG.Database.Port, global.GO_CONFIG.Database.Password))
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Postgresql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormSet(db *gorm.DB) {
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return global.GO_CONFIG.Database.TablePrefix + defaultTableName
	}

	// 设置日志
	if global.GO_CONFIG.App.RunMode == "debug" {
		db.LogMode(true)
	}

	// 设置迁移
	db.AutoMigrate(
		basic.Vendor{},
		model.User{},
	)

	// 设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量
	db.DB().SetMaxOpenConns(100)
}
