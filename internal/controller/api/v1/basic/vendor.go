/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 23:19:52
 * @FilePath: /potato/internal/controller/api/v1/basic/vendor.go
 */
package basic

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode"
	"github.com/viletyy/yolk/convert"
)

type Vendor struct{}

func NewVendor() Vendor {
	return Vendor{}
}

// @Summary 系统厂商列表
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param token header string true "auth by /auth"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param name query string false "系统厂商名称" maxlength(100)
// @Param uuid query int false "系统厂商云id"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors [get]
func (vendor Vendor) List(c *gin.Context) {
	param := service.VendorListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	paper := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	total, err := svc.CountVendor(&service.CountVendorRequest{Name: param.Name, Uuid: param.Uuid})
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.CountVendor err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountVendorFail)
		return
	}
	vendors, err := svc.GetVendorList(&param, &paper)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.GetVendorList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetVendorListFail)
		return
	}

	response.ToResponseList(vendors, total)
}

// @Summary 系统厂商
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param token header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors/{id} [get]
func (vendor Vendor) Get(c *gin.Context) {
	vendorId, err := convert.StrTo(c.Param("id")).Int64()
	response := app.NewResponse(c)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	param := service.VendorRequest{ID: vendorId}

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	dbVendor, err := svc.GetVendor(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.GetVendor err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetVendorFail)
		return
	}

	response.ToResponse(dbVendor)
}

// @Summary 新增系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param token header string true "auth by /auth"
// @Param name formData string true "系统厂商名称" minlength(1) maxlength(100)
// @Param uuid formData int false "系统厂商云id"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors [post]
func (vendor Vendor) Create(c *gin.Context) {
	param := service.CreateVendorRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs.Errors())
		return
	}

	svc := service.New(c.Request.Context())
	dbVendor, err := svc.CreateVendor(&param)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("svc.CreateVendor err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateVendorFail)
		return
	}

	response.ToResponse(dbVendor)
}

// @Summary 修改系统厂商
// @Tags vendors
// @Description
// @Accept mpfd
// @Produce json
// @Param token header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Param name formData string false "系统厂商名称" minlength(1) maxlength(100)
// @Param uuid formData int false "系统厂商云id"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors/{id} [patch]
func (vendor Vendor) Update(c *gin.Context) {
	vendorId, err := convert.StrTo(c.Param("id")).Int64()
	response := app.NewResponse(c)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	param := service.UpdateVendorRequest{ID: vendorId}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.GO_LOG.Sugar().Errorf("app.BindAndValid errs: %v", errs)
		response.ToResponseErrors(errs)
		return
	}
	svc := service.New(c.Request.Context())
	dbVendor, err := svc.UpdateVendor(&param)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GO_LOG.Sugar().Errorf("svc.UpdateVendor err: %v", err)
			response.ToErrorResponse(errcode.ErrorGetVendorFail)
			return
		} else {
			global.GO_LOG.Sugar().Errorf("svc.UpdateVendor err: %v", err)
			response.ToErrorResponse(errcode.ErrorUpdateVendorFail)
			return
		}
	}

	response.ToResponse(dbVendor)
}

// @Summary 删除系统厂商
// @Tags vendors
// @Description
// @Accept json
// @Produce json
// @Param token header string true "auth by /auth"
// @Param id path int true "系统厂商 ID"
// @Success 200 {object} basic.Vendor "请求成功"
// @Router /v1/vendors/{id} [delete]
func (vendor Vendor) Delete(c *gin.Context) {
	vendorId, err := convert.StrTo(c.Param("id")).Int64()
	response := app.NewResponse(c)
	if err != nil {
		global.GO_LOG.Sugar().Errorf("convert.StrTo err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	param := service.DeleteVendorRequest{ID: vendorId}
	svc := service.New(c.Request.Context())
	dbvendor, err := svc.DeleteVendor(&param)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GO_LOG.Sugar().Errorf("svc.DeleteVendor err: %v", err)
			response.ToErrorResponse(errcode.ErrorGetVendorFail)
			return
		} else {
			global.GO_LOG.Sugar().Errorf("svc.DeleteVendor err: %v", err)
			response.ToErrorResponse(errcode.ErrorDeleteVendorFail)
			return
		}
	}

	errcode.Success.Data = dbvendor
	response.ToErrorResponse(errcode.Success)
}
