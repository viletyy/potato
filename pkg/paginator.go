/*
 * @Date: 2021-06-10 15:27:36
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 15:52:48
 * @FilePath: /potato/pkg/paginator.go
 */
package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/yolk/convert"
)

func GetPageInfo(c *gin.Context) (page, pageSize int) {
	page, _ = convert.StrTo(c.DefaultQuery("page", "1")).Int()
	pageSize, _ = convert.StrTo(c.DefaultQuery("page_size", "10")).Int()
	return
}
