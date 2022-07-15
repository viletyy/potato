/*
 * @Date: 2021-03-22 09:42:09
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 23:53:51
 * @FilePath: /potato/global/global.go
 */
package global

import (
	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/viletyy/potato/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GO_DB     *gorm.DB
	GO_REDIS  *redis.Client
	GO_CONFIG *config.Config
	GO_VP     *viper.Viper
	GO_LOG    *zap.Logger
	GO_TRACER opentracing.Tracer
)
