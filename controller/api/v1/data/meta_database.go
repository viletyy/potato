package data

import (
	"github.com/gin-gonic/gin"
	_ "github.com/viletyy/potato/models/data"
	data2 "github.com/viletyy/potato/models/data"
	"github.com/viletyy/potato/pkg/e"
	"github.com/viletyy/potato/pkg/setting"
	"github.com/viletyy/potato/pkg/util"
	"net/http"
)

// @Summary 数据源列表
// @Tags meta_databases
// @Description
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code" : 200, "data" : {}, "msg": "ok" }"
// @Router /v1/meta_databases [get]
func GetMetaDatabases(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	data["lists"] = data2.GetMetaDatabases(util.GetPage(c), setting.PageSize, maps)
	data["total"] = data2.GetMetaDatabaseTotal(maps)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}