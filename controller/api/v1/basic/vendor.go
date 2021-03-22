package basic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/models/basic"
	"github.com/viletyy/potato/utils"
	"go.uber.org/zap"
)

// @Summary 系统厂商列表
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param data body models.VendorSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/vendors [get]
func GetVendors(c *gin.Context) {
	var search basic.VendorSearch

	if err := c.ShouldBindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if result, err := basic.GetVendors(&search); err != nil {
		global.GO_LOG.Error("获取失败!", zap.Any("err", err))
		utils.FailWithMessage("获取失败", c)
	} else {
		utils.OkWithDetailed(result, "获取成功", c)
	}
}

// @Summary 新增系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param data body basic.Vendor true "Vendor模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/vendors [post]
func AddVendor(c *gin.Context) {
	var vendor basic.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := basic.CreateVendor(vendor); err != nil {
		global.GO_LOG.Error("创建失败!", zap.Any("err", err))
		utils.FailWithMessage("创建失败", c)
	} else {
		utils.OkWithMessage("创建成功", c)
	}
}

// @Summary 修改系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Param data body basic.Vendor true "Vendor模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/vendors/{id} [patch]
func UpdateVendor(c *gin.Context) {
	var vendor basic.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := basic.UpdateVendor(&vendor); err != nil {
		global.GO_LOG.Error("更新失败!", zap.Any("err", err))
		utils.FailWithMessage("更新失败", c)
	} else {
		utils.OkWithMessage("更新成功", c)
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
	var vendor basic.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := basic.DeleteVendor(&vendor); err != nil {
		global.GO_LOG.Error("删除失败!", zap.Any("err", err))
		utils.FailWithMessage("删除失败", c)
	} else {
		utils.OkWithMessage("删除成功", c)
	}
}
