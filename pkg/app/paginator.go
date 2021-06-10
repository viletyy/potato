/*
 * @Date: 2021-06-10 15:27:36
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 16:58:28
 * @FilePath: /potato/pkg/app/paginator.go
 */
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/yolk/convert"
)

func GetPageInfo(c *gin.Context) (page, pageSize int) {
	page, _ = convert.StrTo(c.DefaultQuery("page", "1")).Int()
	pageSize, _ = convert.StrTo(c.DefaultQuery("page_size", "10")).Int()
	return
}

func GetPage(c *gin.Context) int {
	page, _ := convert.StrTo(c.Query("page")).Int()
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize, _ := convert.StrTo(c.Query("page_size")).Int()
	if pageSize <= 0 {
		return int(global.GO_CONFIG.App.PageSize)
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
