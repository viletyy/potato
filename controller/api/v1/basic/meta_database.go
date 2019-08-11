package basic

import (
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/models/basic"
	_ "github.com/viletyy/potato/models/basic"
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
// @Success 200 {string} json "{"code" : 200, "basic" : {}, "msg": "ok" }"
// @Router /v1/meta_databases [get]
func GetMetaDatabases(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	data["lists"] = basic.GetMetaDatabases(util.GetPage(c), setting.PageSize, maps)
	data["total"] = basic.GetMetaDatabaseTotal(maps)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"basic": data,
	})
}