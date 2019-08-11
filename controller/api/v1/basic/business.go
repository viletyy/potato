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

// @Summary 业务系统列表
// @Tags businesses
// @Description
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code" : 200, "data" : {}, "msg" : "ok"}"
// @Router /v1/businesses [get]
func GetBusinesses(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	data["lists"] = basic.GetBusinesses(util.GetPage(c), setting.PageSize, maps)
	data["total"] = basic.GetBusinessTotal(maps)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary 新增业务系统
// @Tags businesses
// @Description
// @Accept mpfd
// @Produce json
// @Param name formData string true "业务系统 名称"
// @Param desc formData string false "业务系统 描述"
// @Param c_id formData int false "业务系统 云端id"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/businesses [post]
func AddBusiness(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	cId := com.StrTo(c.PostForm("c_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if ! basic.ExistVendorByName(name) {
			code = e.SUCCESS
			basic.AddBusiness(name, desc, cId)
		} else {
			code = e.ERROR_EXIST_BUSINESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

// @Summary 修改业务系统
// @Tags businesses
// @Description
// @Accept mpfd
// @Produce json
// @Param id path int true "业务系统 ID"
// @Param name formData string false "业务系统 名称"
// @Param desc formData string false "业务系统 描述"
// @Param c_id formData string false "业务系统 云端id"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/businesses/{id} [patch]
func EditBusiness(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	desc := c.PostForm("desc")
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
			if desc != "" {
				data["desc"] = desc
			}
			if cId != 0 {
				data["c_id"] = cId
			}
			basic.EditBusiness(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_BUSINESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary 删除业务系统
// @Tags businesses
// @Description
// @Accept json
// @Produce json
// @Param id path int true "业务系统 ID"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/businesses/{id} [delete]
func DeleteBusiness(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if basic.ExistBusinessById(id) {
			basic.DeleteBusiness(id)
		} else {
			code = e.ERROR_NOT_EXIST_BUSINESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}