package basic

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/models/basic"
	"github.com/viletyy/potato/pkg/e"
	"github.com/viletyy/potato/pkg/setting"
	"github.com/viletyy/potato/pkg/util"
	"net/http"
)

// @Summary 系统厂商列表
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Success 200 {string} json "{"code" : 200, "data" : {}, "msg" : "ok"}"
// @Router /v1/vendors [get]
func GetVendors(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	data["lists"] = basic.GetVendors(util.GetPage(c), setting.PageSize, maps)
	data["total"] = basic.GetVendorsTotal(maps)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param name formData string true "系统厂商 名称"
// @Param c_id formData int false "系统厂商 云端id"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/vendors [post]
func AddVendor(c *gin.Context) {
	name := c.PostForm("name")
	cId := com.StrTo(c.PostForm("c_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if ! basic.ExistVendorByName(name) {
			code = e.SUCCESS
			basic.AddVendor(name, cId)
		} else {
			code = e.ERROR_EXIST_VENDOR
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

// @Summary 修改系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Param name formData string false "系统厂商 名称"
// @Param c_id formData int false "系统厂商 云端id"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/vendors/{id} [patch]
func EditVendor(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	name := c.PostForm("name")
	cId := com.StrTo(c.PostForm("c_id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID必须大于0")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if basic.ExistVendorById(id) {
			if name != "" {
				data["name"] = name
			}
			if cId > 0 {
				data["c_id"] = cId
			}
			basic.EditVendor(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_VENDOR
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})

}

// @Summary 删除系统厂商
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/vendors/{id} [delete]
func DeleteVendor(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if basic.ExistVendorById(id) {
			basic.DeleteVendor(id)
		} else {
			code = e.ERROR_NOT_EXIST_VENDOR
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
	})
}