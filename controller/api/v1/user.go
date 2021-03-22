/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-23 00:44:10
 * @FilePath: /potato/controller/api/v1/user.go
 */
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/models"
	"github.com/viletyy/potato/utils"
	"go.uber.org/zap"
)

type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// @Summary 用户验证
// @Description
// @Accept json
// @Produce json
// @Param data body models.User true "User模型"
// @Success 200 {string} json "{"code" : 200, "data" : {"token" : ""}, "msg" : "ok"}"
// @Router /v1/auth [get]
func GetUserAuth(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	mUser, gErr := models.GetUserByUsername(user.Username)
	if gErr != nil {
		global.GO_LOG.Error("查找用户失败", zap.Any("err", gErr))
		utils.FailWithMessage("查找用户失败", c)
	}

	isTrue := mUser.CheckPassword(user.Password)
	if !isTrue {
		global.GO_LOG.Error("用户密码不正确")
		utils.FailWithMessage("用户密码不正确", c)
	}

	token, tokenErr := utils.GenerateToken(mUser.ID)
	if tokenErr != nil {
		global.GO_LOG.Error("获取token失败", zap.Any("err", tokenErr))
		utils.FailWithMessage("获取token失败", c)
	}

	utils.OkWithDetailed(LoginResponse{
		User:  mUser,
		Token: token,
	}, "登录成功", c)
}

// @Summary 新增用户
// @Tags users
// @Description
// @Accept mpfd
// @Produce json
// @Param data body basic.User true "User模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/users [post]
func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := models.CreateUser(user); err != nil {
		global.GO_LOG.Error("创建失败!", zap.Any("err", err))
		utils.FailWithMessage("创建失败", c)
	} else {
		utils.OkWithMessage("创建成功", c)
	}
}
