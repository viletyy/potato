/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 17:56:42
 * @FilePath: /potato/internal/controller/api/v1/basic/vendor.go
 */
package basic

import (
	"github.com/gin-gonic/gin"
)

type Vendor struct{}

// @Summary 系统厂商列表
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param name query string false "名称" maxlength(100)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors [get]
func (vendor Vendor) List(c *gin.Context) {}

// @Summary 新增系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors [post]
func (vendor Vendor) Create(c *gin.Context) {}

// @Summary 修改系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors/{id} [patch]
func (vendor Vendor) Update(c *gin.Context) {}

// @Summary 删除系统厂商
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param Authorization header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors/{id} [delete]
func (vendor Vendor) Delete(c *gin.Context) {
}
