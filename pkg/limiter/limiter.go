/*
 * @Date: 2021-06-13 22:06:52
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-13 22:13:35
 * @FilePath: /potato/pkg/limiter/limiter.go
 */
package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterInterface interface {
	Key(c *gin.Context) string                              // 获取对应的限流器的键值对名称。
	GetBucket(key string) (*ratelimit.Bucket, bool)         // 获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterInterface // 新增多个令牌桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration // 间隔多久时间放N个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}
