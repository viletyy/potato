/*
 * @Date: 2021-07-08 13:54:38
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-08 14:01:18
 * @FilePath: /potato/initialize/cron.go
 */
package initialize

import (
	"github.com/robfig/cron/v3"
	"github.com/viletyy/potato/internal/job"
)

func Cron() {
	c := cron.New()

	c.AddJob("* * * * *", job.TestJob{})

	c.Start()

	select {}
}
