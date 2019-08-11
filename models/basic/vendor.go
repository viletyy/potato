package basic

import (
	"github.com/viletyy/potato/pkg/util"
)

type Vendor struct {
	util.Model

	Name string `json:"name"`
	CId int `json:"c_id"`
}

func GetVendors(pageNum int, pageSize int, maps interface{}) (vendors []Vendor) {
	util.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&vendors)
	return
}

func GetVendorsTotal(maps interface{}) (count int) {
	util.DB.Model(&Vendor{}).Where(maps).Count(&count)
	return
}

func ExistVendorByName(name string) bool {
	var vendor Vendor
	util.DB.Select("id").Where("name = ?", name).First(&vendor)
	if vendor.ID > 0 {
		return true
	}
	return false
}

func ExistVendorById(id int) bool {
	var vendor Vendor
	util.DB.Select("id").Where("id = ?", id).First(&vendor)
	if vendor.ID > 0 {
		return true
	}
	return false
}

func AddVendor(name string, cId int) bool {
	util.DB.Create(&Vendor {
		Name: name,
		CId: cId,
	})

	return true
}

func EditVendor(id int, data interface{}) bool {
	util.DB.Model(&Vendor{}).Where("id = ?", id).Update(data)

	return true
}

func DeleteVendor(id int) bool {
	util.DB.Where("id = ?", id).Delete(&Vendor{})

	return true
}