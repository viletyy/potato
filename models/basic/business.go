/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 23:54:52
 * @FilePath: /potato/models/basic/business.go
 */
package basic

import (
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/utils"
)

type BusinessSearch struct {
	Business
	utils.PageInfo
}

type Business struct {
	global.Model

	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Uuid        int    `json:"uuid"`
}

func GetBusinesses(search *BusinessSearch) (searchResult utils.SearchResult, err error) {
	var businesses []Business
	offset := search.PageInfo.PageSize * (search.PageInfo.Page - 1)
	limit := search.PageInfo.Page
	db := global.GO_DB.Where(search.Business)
	err = db.Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&businesses).Error
	if err != nil {
		return
	}
	searchResult.Page = search.PageInfo.Page
	searchResult.PageSize = search.PageInfo.PageSize
	searchResult.List = businesses
	return
}

func GetBusinessById(id int) (business Business, err error) {
	err = global.GO_DB.Where("id = ?", id).First(&business).Error
	return
}

func GetBusinessByName(name string) (business Business, err error) {
	err = global.GO_DB.Where("name = ?", name).First(&business).Error
	return
}

func GetBusinessByUuid(uuid int64) (business Business, err error) {
	err = global.GO_DB.Where("uuid = ?", uuid).First(&business).Error
	return
}

func ExistBusinessById(id int) bool {
	var business Business
	global.GO_DB.Where("id = ?", id).First(&business)

	return business.ID > 0
}

func ExistBusinessByName(name string) bool {
	var business Business
	global.GO_DB.Where("name = ?", name).First(&business)

	return business.ID > 0
}

func ExistBusinessByUuid(uuid int64) bool {
	var business Business
	global.GO_DB.Where("uuid = ?", uuid).First(&business)

	return business.ID > 0
}

func CreateBusiness(business Business) (err error) {
	err = global.GO_DB.Create(&business).Error

	return err
}

func UpdateBusiness(business *Business) (err error) {
	err = global.GO_DB.Save(&business).Error
	return
}

func DeleteBusiness(business *Business) (err error) {
	err = global.GO_DB.Delete(&business).Error
	return
}
