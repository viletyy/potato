/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 15:26:27
 * @FilePath: /potato/internal/model/basic/meta_database.go
 */
package basic

import (
	"time"

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/model"
	"github.com/viletyy/potato/pkg"
)

type MetaDatabaseSearch struct {
	MetaDatabase
	pkg.PageInfo
}

type MetaDatabase struct {
	model.Model
	Name            string    `json:"name"`
	Adapter         string    `json:"adapter"`
	Host            string    `json:"host"`
	Port            string    `json:"port"`
	DbName          string    `json:"db_name"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	LastConnectTime time.Time `json:"last_connect_time"`
	Usable          bool      `json:"usable"`
	Comment         string    `json:"comment"`
	VendorId        int64     `json:"vendor_id"`
	BusinessId      int64     `json:"business_id"`
	Vendor          Vendor
	Business        Business
}

func GetMetaDatabases(search *MetaDatabaseSearch) (searchResult pkg.SearchResult, err error) {
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
