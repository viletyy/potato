/*
 * @Date: 2021-06-10 17:57:48
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 23:19:35
 * @FilePath: /potato/internal/service/vendor.go
 */
package service

import (
	"github.com/viletyy/potato/internal/model/basic"
	"github.com/viletyy/potato/pkg/app"
)

type CountVendorRequest struct {
	Name string `form:"name" json:"name" validate:"max=100"`
	Uuid int    `form:"uuid" json:"uuid" `
}

type VendorListRequest struct {
	Name string `form:"name" json:"name" validate:"max=100"`
	Uuid int    `form:"uuid" json:"uuid"`
}

type VendorRequest struct {
	ID int64 `form:"id" json:"id" validate:"required,gte=1"`
}

type CreateVendorRequest struct {
	Name string `form:"name" json:"name" validate:"required"`
	Uuid int    `form:"uuid"  json:"uuid"`
}

type UpdateVendorRequest struct {
	ID   int64  `form:"id" json:"id" validate:"required,gte=1"`
	Name string `form:"name" json:"name"`
	Uuid int    `form:"uuid" json:"uuid"`
}

type DeleteVendorRequest struct {
	ID int64 `json:"id" validate:"required,gte=1"`
}

func (svc *Service) CountVendor(param *CountVendorRequest) (int, error) {
	return svc.dao.CountVendor(param.Name, param.Uuid)
}

func (svc *Service) GetVendorList(param *VendorListRequest, pager *app.Pager) ([]basic.Vendor, error) {
	return svc.dao.GetVendorList(param.Name, param.Uuid, pager.Page, pager.PageSize)
}

func (svc *Service) GetVendor(param *VendorRequest) (basic.Vendor, error) {
	return svc.dao.GetVendor(param.ID)
}

func (svc *Service) CreateVendor(param *CreateVendorRequest) (basic.Vendor, error) {
	return svc.dao.CreateVendor(param.Name, param.Uuid)
}

func (svc *Service) UpdateVendor(param *UpdateVendorRequest) (basic.Vendor, error) {
	return svc.dao.UpdateVendor(param.ID, param.Name, param.Uuid)
}

func (svc *Service) DeleteVendor(param *DeleteVendorRequest) (basic.Vendor, error) {
	return svc.dao.DeleteVendor(param.ID)
}
