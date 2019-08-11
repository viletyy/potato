package basic

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/pkg/e"
	"net/http"
)

// @Summary 元数据列表
// @Tags meta_tables
// @Description
// @Accept json
// @Produce json
// @Param id path int true "数据源 ID"
// @Success 200 {string} json "{"code" : 200, "basic" : {}, "msg": "ok" }"
// @Router /v1/meta_databases/{id}/meta_tables [get]
func GetMetaTables(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"basic": data,
	})
}