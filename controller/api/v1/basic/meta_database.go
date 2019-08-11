package basic

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
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

// @Summary 新增数据源
// @Tags meta_databases
// @Description
// @Accept mpfd
// @Produce json
// @Param name formData string true "数据源 名称"
// @Param host formData string true "数据源 地址"
// @Param port formData int true "数据源 端口号"
// @Param db_name formData string true "数据源 数据库名称"
// @Param username formData string true "数据源 用户名"
// @Param password formData string true "数据源 密码"
// @Param comment formData string false "数据源 备注"
// @Param vendor_id formData int true "系统厂商 id"
// @Param business_id formData int true "业务系统 id"
// @Success 200 {string} json "{"code": 200, data: {}, "msg" : "ok"}"
// @Router /v1/meta_databases [post]
func AddMetaDatabase(c *gin.Context) {
	name := c.PostForm("name")
	host := c.PostForm("host")
	port := com.StrTo(c.PostForm("port")).MustInt()
	dbName := c.PostForm("db_name")
	username := c.PostForm("username")
	password := c.PostForm("password")
	comment := c.PostForm("comment")
	vendorId := com.StrTo(c.PostForm("vendor_id")).MustInt()
	businessId := com.StrTo(c.PostForm("business_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(host, "host").Message("地址不能为空")
	valid.Required(port, "port").Message("端口号不能为空")
	valid.Required(dbName, "db_name").Message("数据库名称不能为空")
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Min(vendorId, 1,"vendor_id").Message("必须是有效的系统厂商ID")
	valid.Min(businessId, 1,"business_id").Message("必须是有效的业务系统ID")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ! basic.ExistVendorById(vendorId) {
		code = e.ERROR_NOT_EXIST_VENDOR
	}

	if ! basic.ExistBusinessById(businessId) {
		code = e.ERROR_NOT_EXIST_BUSINESS
	}

	if ! valid.HasErrors() {
		if ! basic.ExistMetaDatabaseByName(name) {
			data["Name"] = name
			data["Host"] = host
			data["Port"] = port
			data["DbName"] = dbName
			data["Username"] = username
			data["Password"] = password
			data["Comment"] = comment
			data["VendorId"] = vendorId
			data["BusinessId"] = businessId
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_META_DATABASE
		}
	}

	if code == 200 {
		basic.AddMetaDatabase(data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})

}

// @Summary 修改数据源
// @Tags meta_databases
// @Description
// @Accept mpfd
// @Produce json
// @Param id path int true "数据源 ID"
// @Param name formData string false "数据源 名称"
// @Param host formData string false "数据源 地址"
// @Param port formData int false "数据源 端口号"
// @Param db_name formData string false "数据源 数据库名称"
// @Param username formData string false "数据源 用户名"
// @Param password formData string false "数据源 密码"
// @Param comment formData string false "数据源 备注"
// @Param vendor_id formData int false "系统厂商 id"
// @Param business_id formData int false "业务系统 id"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/meta_databases/{id} [patch]
func EditMetaDatabase(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	host := c.PostForm("host")
	port := com.StrTo(c.PostForm("port")).MustInt()
	dbName := c.PostForm("db_name")
	username := c.PostForm("username")
	password := c.PostForm("password")
	comment := c.PostForm("comment")
	vendorId := com.StrTo(c.PostForm("vendor_id")).MustInt()
	businessId := com.StrTo(c.PostForm("business_id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("必须是有效的数据源id")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if basic.ExistMetaDatabaseById(id) {
			code = e.SUCCESS
			if name != "" {
				data["name"] = name
			}
			if host != "" {
				data["host"] = host
			}
			if port != 0 {
				data["port"] = port
			}
			if dbName != "" {
				data["db_name"] = dbName
			}
			if username != "" {
				data["username"] = username
			}
			if password != "" {
				data["password"] = password
			}
			if comment != "" {
				data["comment"] = comment
			}
			if vendorId != 0 {
				if basic.ExistVendorById(vendorId) {
					data["vendor_id"] = vendorId
				} else {
					code = e.ERROR_NOT_EXIST_VENDOR
				}
			}
			if businessId != 0 {
				if basic.ExistBusinessById(businessId) {
					data["business_id"] = businessId
				} else {
					code = e.ERROR_NOT_EXIST_BUSINESS
				}
			}
		} else {
			code = e.ERROR_NOT_EXIST_META_DATABASE
		}
	}
	if code == 200 {
		basic.EditMetaDatabase(id, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary 删除数据源
// @Tags meta_databases
// @Description
// @Accept json
// @Produce json
// @Param id path int true "数据源 ID"
// @Success 200 {string} json "{"code" : 200, "msg" : "ok"}"
// @Router /v1/meta_databases/{id} [delete]
func DeleteMetaDatabase(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("必须是有效的数据源ID")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if basic.ExistMetaDatabaseById(id) {
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_META_DATABASE
		}
	}

	if code == 200 {
		basic.DeleteMetaDatabase(id)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
	})
}