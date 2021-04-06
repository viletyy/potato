/*
 * @Date: 2021-03-22 09:42:09
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-06 18:18:15
 * @FilePath: /potato/global/global.go
 */
package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/viletyy/potato/config"
	"go.uber.org/zap"
)

var (
	GO_DB     *gorm.DB
	GO_REDIS  *redis.Client
	GO_CONFIG *config.Config
	GO_VP     *viper.Viper
	GO_LOG    *zap.Logger
)
