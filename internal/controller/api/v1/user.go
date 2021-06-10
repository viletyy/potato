/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 15:23:41
 * @FilePath: /potato/internal/controller/api/v1/user.go
 */
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/model"
	"github.com/viletyy/potato/pkg"
	"github.com/viletyy/yolk/crypt"
	"go.uber.org/zap"
)

type AuthResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=6"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=6"`
	Nickname string `json:"nickname"`
}

// @Summary 用户验证
// @Description
// @Accept json
// @Produce json
// @Param data body AuthRequest true "用户名,密码"
// @Success 200 {string} json "{"code" : 200, "data" : {"token" : ""}, "msg" : "ok"}"
// @Router /v1/auth [post]
func Auth(c *gin.Context) {
	var user AuthRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": errs.Translate(pkg.Trans),
			})
			return
		}
	}

	mUser, gErr := model.GetUserByUsername(user.Username)
	if gErr != nil {
		global.GO_LOG.Error("查找用户失败", zap.Any("err", gErr))
		pkg.FailWithMessage("查找用户失败", c)
		return
	}

	isTrue := mUser.CheckPassword(user.Password)
	if !isTrue {
		global.GO_LOG.Error("用户密码不正确")
		pkg.FailWithMessage("用户密码不正确", c)
		return
	}

	token, tokenErr := pkg.GenerateToken(mUser.ID)
	if tokenErr != nil {
		global.GO_LOG.Error("获取token失败", zap.Any("err", tokenErr))
		pkg.FailWithMessage("获取token失败", c)
		return
	}

	pkg.OkWithDetailed(AuthResponse{
		User:  mUser,
		Token: token,
	}, "登录成功", c)
}

// @Summary 注册用户
// @Description
// @Accept json
// @Produce json
// @Param data body RegisterRequest true "用户名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/register [post]
func Register(c *gin.Context) {
	var user RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": errs.Translate(pkg.Trans),
			})
			return
		}
	}

	if isExsit := model.ExistUserByUsername(user.Username); isExsit {
		global.GO_LOG.Error("用户已存在")
		pkg.FailWithMessage("用户已存在", c)
		return
	}

	if err := model.CreateUser(model.User{Username: user.Username, Password: crypt.Md5Encode(user.Password), Nickname: user.Nickname}); err != nil {
		global.GO_LOG.Error("创建失败!", zap.Any("err", err))
		pkg.FailWithMessage("创建失败", c)
	} else {
		pkg.OkWithMessage("创建成功", c)
	}
}
