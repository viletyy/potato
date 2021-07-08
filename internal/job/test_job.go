/*
 * @Date: 2021-07-08 14:01:46
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-08 14:02:36
 * @FilePath: /potato/internal/job/test_job.go
 */
package job

import "fmt"

type TestJob struct{}

func (t TestJob) Run() {
	fmt.Println("i'm test job")
}
