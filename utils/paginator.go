/*
 * @Date: 2021-03-24 10:12:24
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-24 10:30:58
 * @FilePath: /potato/utils/paginator.go
 */
package utils

import "github.com/gin-gonic/gin"

func GetPageInfo(c *gin.Context) (page, pageSize int) {
	page = ToInt(c.DefaultQuery("page", "1"))
	pageSize = ToInt(c.DefaultQuery("page_size", "10"))
	return
}
