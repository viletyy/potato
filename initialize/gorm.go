/*
 * @Date: 2021-03-22 10:12:38
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 16:56:26
 * @FilePath: /potato/initialize/gorm.go
 */
package initialize

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/viletyy/potato/global"
)

var dbConfig = global.GO_CONFIG.Database

func Gorm() *gorm.DB {
	switch dbConfig.Type {
	case "mysql":
		return GormMysql()
	case "postgresql":
		return GormPostgresql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Mysql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormPostgresql() *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable password=%s", dbConfig.Host, dbConfig.User, dbConfig.Name, dbConfig.Port, dbConfig.Password))
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Postgresql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormSet(db *gorm.DB) {
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbConfig.TablePrefix + defaultTableName
	}

	// 设置日志
	db.LogMode(true)

	// 设置迁移
	db.AutoMigrate()

	// 设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量
	db.DB().SetMaxOpenConns(100)
}
