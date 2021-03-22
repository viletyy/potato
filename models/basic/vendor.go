/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 23:52:13
 * @FilePath: /potato/models/basic/vendor.go
 */
package basic

import (
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/utils"
)

type VendorSearch struct {
	Vendor
	utils.PageInfo
}

type Vendor struct {
	global.Model

	Name string `json:"name" binding:"require"`
	Uuid int    `json:"uuid"`
}

func GetVendors(search *VendorSearch) (searchResult utils.SearchResult, err error) {
	var vendors []Vendor
	offset := search.PageInfo.PageSize * (search.PageInfo.Page - 1)
	limit := search.PageInfo.Page
	db := global.GO_DB.Where(search.Vendor)
	err = db.Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&vendors).Error
	if err != nil {
		return
	}
	searchResult.Page = search.PageInfo.Page
	searchResult.PageSize = search.PageInfo.PageSize
	searchResult.List = vendors
	return
}

func GetVendorById(id int) (vendor Vendor, err error) {
	err = global.GO_DB.Where("id = ?", id).First(&vendor).Error
	return
}

func GetVendorByName(name string) (vendor Vendor, err error) {
	err = global.GO_DB.Where("name = ?", name).First(&vendor).Error
	return
}

func GetVendorByUuid(uuid int64) (vendor Vendor, err error) {
	err = global.GO_DB.Where("uuid = ?", uuid).First(&vendor).Error
	return
}

func ExistVendorById(id int) bool {
	var vendor Vendor
	global.GO_DB.Where("id = ?", id).First(&vendor)

	return vendor.ID > 0
}

func ExistVendorByName(name string) bool {
	var vendor Vendor
	global.GO_DB.Where("name = ?", name).First(&vendor)

	return vendor.ID > 0
}

func ExistVendorByUuid(uuid int64) bool {
	var vendor Vendor
	global.GO_DB.Where("uuid = ?", uuid).First(&vendor)

	return vendor.ID > 0
}

func CreateVendor(vendor Vendor) (err error) {
	err = global.GO_DB.Create(&vendor).Error

	return err
}

func UpdateVendor(vendor *Vendor) (err error) {
	err = global.GO_DB.Save(&vendor).Error
	return
}

func DeleteVendor(vendor *Vendor) (err error) {
	err = global.GO_DB.Delete(&vendor).Error
	return
}
