/*
 * @Date: 2021-03-22 10:12:42
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 17:02:17
 * @FilePath: /potato/initialize/redis.go
 */
package initialize

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/viletyy/potato/global"
)

var redisConfig = global.GO_CONFIG.Redis

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       int(redisConfig.Db),
	})

	RedisSet(rdb)
	return rdb
}

func RedisSet(rdb *redis.Client) {
	_, pingErr := rdb.Ping().Result()
	if pingErr != nil {
		global.GO_LOG.Error(fmt.Sprintf("Redis Connection Error: %v", pingErr))
	}
}
