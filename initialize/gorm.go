/*
 * @Date: 2021-03-22 10:12:38
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:27:45
 * @FilePath: /potato/initialize/gorm.go
 */
package initialize

import (
	"fmt"

	"github.com/viletyy/potato/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.GO_CONFIG.Database.User, global.GO_CONFIG.Database.Password, global.GO_CONFIG.Database.Host, global.GO_CONFIG.Database.Port, global.GO_CONFIG.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Mysql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormPostgresql() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", global.GO_CONFIG.Database.Host, global.GO_CONFIG.Database.User, global.GO_CONFIG.Database.Password, global.GO_CONFIG.Database.Name, global.GO_CONFIG.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Postgresql Gorm Open Error: %v", err))
	}
	GormSet(db)
	return db
}

func GormSet(db *gorm.DB) {
	if global.GO_CONFIG.App.RunMode != "debug" {
		logger := zapgorm2.New(global.GO_LOG)
		logger.SetAsDefault()
		logger.LogLevel = gormlogger.Info
		db.Logger = logger
	}

	sqlDB, err := db.DB()
	if err != nil {
		global.GO_LOG.Error(fmt.Sprintf("Gorm setting db.DB(): %v ", err))
	}

	// 设置空闲连接池中的最大连接数
	sqlDB.SetConnMaxIdleTime(10)

	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
}
