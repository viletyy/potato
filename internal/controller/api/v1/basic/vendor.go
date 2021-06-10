package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/model/basic"
	"github.com/viletyy/potato/pkg"
	"go.uber.org/zap"
)

type CreateVendorRequest struct {
	Name string `json:"name" validate:"required"`
	Uuid int    `json:"uuid"`
}

type UpdateVendorRequest struct {
	Name string `json:"name"`
	Uuid int    `json:"uuid"`
}

// @Summary 系统厂商列表
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param name query string false "系统厂商名称"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/vendors [get]
func GetVendors(c *gin.Context) {
	var search basic.VendorSearch

	search.Name = c.Query("name")
	search.Page, search.PageSize = pkg.GetPageInfo(c)

	if result, err := basic.GetVendors(&search); err != nil {
		global.GO_LOG.Error("获取失败!", zap.Any("err", err))
		pkg.FailWithMessage("获取失败", c)
	} else {
		pkg.OkWithDetailed(result, "获取成功", c)
	}
}

// @Summary 新增系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param data body CreateVendorRequest true "Vendor模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/vendors [post]
func CreateVendor(c *gin.Context) {
	var vendor CreateVendorRequest
	if err := c.ShouldBindJSON(&vendor); err != nil {
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

	if exist := basic.ExistVendorByName(vendor.Name); exist {
		global.GO_LOG.Error("该系统厂商名称已存在!")
		pkg.FailWithMessage("该系统厂商名称已存在", c)
		return
	}
	if exist := basic.ExistVendorByUuid(vendor.Uuid); exist {
		global.GO_LOG.Error("该系统厂商uuid已存在!")
		pkg.FailWithMessage("该系统厂商uuid已存在", c)
		return
	}

	if mVendor, err := basic.CreateVendor(basic.Vendor{Name: vendor.Name, Uuid: vendor.Uuid}); err != nil {
		global.GO_LOG.Error("创建失败!", zap.Any("err", err))
		pkg.FailWithMessage("创建失败", c)
	} else {
		pkg.OkWithDetailed(mVendor, "创建成功", c)
	}
}

// @Summary 修改系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Param data body UpdateVendorRequest true "Vendor模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/vendors/{id} [patch]
func UpdateVendor(c *gin.Context) {
	vendor, gErr := basic.GetVendorById(c.Param("id"))
	if gErr != nil {
		global.GO_LOG.Error("系统厂商不存在!", zap.Any("err", gErr))
		pkg.FailWithMessage("系统厂商不存在!", c)
		return
	}

	if err := c.ShouldBindJSON(&vendor); err != nil {
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

	if exist := basic.ExistVendorByName(vendor.Name); exist {
		global.GO_LOG.Error("该系统厂商名称已存在!")
		pkg.FailWithMessage("该系统厂商名称已存在", c)
		return
	}
	if exist := basic.ExistVendorByUuid(vendor.Uuid); exist {
		global.GO_LOG.Error("该系统厂商uuid已存在!")
		pkg.FailWithMessage("该系统厂商uuid已存在", c)
		return
	}

	if err := basic.UpdateVendor(&vendor); err != nil {
		global.GO_LOG.Error("更新失败!", zap.Any("err", err))
		pkg.FailWithMessage("更新失败", c)
	} else {
		pkg.OkWithDetailed(vendor, "更新成功", c)
	}
}

// @Summary 删除系统厂商
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/vendors/{id} [delete]
func DeleteVendor(c *gin.Context) {
	vendor, gErr := basic.GetVendorById(c.Param("id"))
	if gErr != nil {
		global.GO_LOG.Error("系统厂商不存在!", zap.Any("err", gErr))
		pkg.FailWithMessage("系统厂商不存在!", c)
		return
	}
	if err := basic.DeleteVendor(&vendor); err != nil {
		global.GO_LOG.Error("删除失败!", zap.Any("err", err))
		pkg.FailWithMessage("删除失败", c)
	} else {
		pkg.OkWithMessage("删除成功", c)
	}
}
