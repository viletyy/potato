/*
 * @Date: 2021-03-22 10:12:42
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-23 09:26:11
 * @FilePath: /potato/initialize/redis.go
 */
package initialize

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/viletyy/potato/global"
)

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.GO_CONFIG.Redis.Host, global.GO_CONFIG.Redis.Port),
		Password: global.GO_CONFIG.Redis.Password,
		DB:       int(global.GO_CONFIG.Redis.Db),
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
