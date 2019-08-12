package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/models"
	"github.com/viletyy/potato/pkg/e"
	"github.com/viletyy/potato/pkg/logging"
	"github.com/viletyy/potato/pkg/setting"
	"github.com/viletyy/potato/pkg/util"
	"net/http"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 用户验证
// @Description
// @Accept json
// @Produce json
// @Param username query string true "用户 用户名"
// @Param password query string true "用户 密码"
// @Success 200 {string} json "{"code" : 200, "data" : {"token" : ""}, "msg" : "ok"}"
// @Router /v1/auth [get]
func GetUserAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := user{Username: username, Password: password}
	ok, err := valid.Valid(&a)

	logging.Info(err)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckUser(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary 用户列表
// @Tags users
// @Description
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code" : 200, "data" : {}, "msg" : "ok"}"
// @Router /v1/users [get]
func GetUsers(c *gin.Context) {
	username := c.Query("username")
	nickname := c.Query("nickname")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if username != "" {
		maps["username"] = username
	}
	if nickname != "" {
		maps["nickname"] = nickname
	}

	data["lists"] = models.GetUsers(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetUsersTotal(maps)
	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary 新增用户
// @Tags users
// @Description
// @Accept mpfd
// @Produce json
// @Param username formData string true "用户 用户名"
// @Param password formData string true "用户 密码"
// @Param nickname formData string true "用户 真实姓名"
// @Success 200 {string} json "{"code" : 200, data: {}, "msg" : "ok"}"
// @Router /v1/users [post]
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Required(nickname, "password").Message("真实姓名不能为空")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if ! models.ExistUserByUsername(username) {
			data["Username"] = username
			data["Password"] = models.GetSecretPassword(password)
			data["Nickname"] = nickname
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_USER
		}
	}

	if code == 200 {
		models.AddUser(data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}