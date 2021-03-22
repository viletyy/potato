/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 23:54:39
 * @FilePath: /potato/models/basic/meta_database.go
 */
package basic

import (
	"time"

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/utils"
)

type MetaDatabaseSearch struct {
	MetaDatabase
	utils.PageInfo
}

type MetaDatabase struct {
	global.Model
	Name            string    `json:"name" binding:"required"`
	Adapter         string    `json:"adapter" binding:"required"`
	Host            string    `json:"host" binding:"required"`
	Port            string    `json:"port" binding:"required"`
	DbName          string    `json:"db_name" binding:"required"`
	Username        string    `json:"username" binding:"required"`
	Password        string    `json:"password" binding:"required"`
	LastConnectTime time.Time `json:"last_connect_time"`
	Usable          bool      `json:"usable"`
	Comment         string    `json:"comment"`
	VendorId        int64     `json:"vendor_id" binding:"required"`
	BusinessId      int64     `json:"business_id" binding:"required"`
	Vendor          Vendor
	Business        Business
}

func GetMetaDatabases(search *MetaDatabaseSearch) (searchResult utils.SearchResult, err error) {
	var metaDatabases []MetaDatabase
	offset := search.PageInfo.PageSize * (search.PageInfo.Page - 1)
	limit := search.PageInfo.Page
	db := global.GO_DB.Where(search.MetaDatabase)
	err = db.Count(&searchResult.Total).Error
	if err != nil {
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&metaDatabases).Error
	if err != nil {
		return
	}
	searchResult.Page = search.PageInfo.Page
	searchResult.PageSize = search.PageInfo.PageSize
	searchResult.List = metaDatabases
	return
}

func GetMetaDatabaseById(id int) (metaDatabase MetaDatabase, err error) {
	err = global.GO_DB.Where("id = ?", id).First(&metaDatabase).Error
	return
}

func GetMetaDatabaseByName(name string) (metaDatabase MetaDatabase, err error) {
	err = global.GO_DB.Where("name = ?", name).First(&metaDatabase).Error
	return
}

func ExistMetaDatabaseById(id int) bool {
	var metaDatabase MetaDatabase
	global.GO_DB.Where("id = ?", id).First(&metaDatabase)

	return metaDatabase.ID > 0
}

func ExistMetaDatabaseByName(name string) bool {
	var metaDatabase MetaDatabase
	global.GO_DB.Where("name = ?", name).First(&metaDatabase)

	return metaDatabase.ID > 0
}

func CreateMetaDatabase(metaDatabase MetaDatabase) (err error) {
	err = global.GO_DB.Create(&metaDatabase).Error

	return err
}

func UpdateMetaDatabase(metaDatabase *MetaDatabase) (err error) {
	err = global.GO_DB.Save(&metaDatabase).Error
	return
}

func DeleteMetaDatabase(metaDatabase *MetaDatabase) (err error) {
	err = global.GO_DB.Delete(&metaDatabase).Error
	return
}
