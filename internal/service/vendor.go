/*
 * @Date: 2021-06-10 17:57:48
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:01:40
 * @FilePath: /potato/internal/service/vendor.go
 */
package service

type CountVendorRequest struct {
	Name string `json:"name" validate:"max=100"`
}

type ListVendorRequest struct {
	Name string `json:"name" validate:"max=100"`
}

type CreateVendorRequest struct {
	Name string `json:"name" validate:"required"`
	Uuid int    `json:"uuid"`
}

type UpdateVendorRequest struct {
	ID   int64  `json:"id" validate:"required,gte=1"`
	Name string `json:"name"`
	Uuid int    `json:"uuid"`
}

type DeleteVendorRequest struct {
	ID int64 `json:"id" validate:"required,gte=1"`
}
