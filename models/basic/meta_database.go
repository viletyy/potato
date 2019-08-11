package basic

import (
	"github.com/viletyy/potato/pkg/util"
	"log"
	"time"
)

type MetaDatabase struct {
	util.Model
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
	DbName string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	LastConnectTime time.Time `json:"last_connect_time"`
	Usable bool `json:"usable"`
	Comment string `json:"comment"`
	Vendor Vendor
	VendorId int `json:"vendor_id"`
	Business Business
	BusinessId int `json:"business_id"`
}

func GetMetaDatabases(pageNum int, pageSize int, maps interface{}) (metaDatabases []MetaDatabase) {
	util.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&metaDatabases)
	return
}

func GetMetaDatabaseTotal(maps interface{}) (count int) {
	util.DB.Model(&MetaDatabase{}).Where(maps).Count(&count)

	return
}

func ExistMetaDatabaseByName(name string) bool {
	var metaDatabase MetaDatabase
	util.DB.Select("id").Where("name = ?", name).First(&metaDatabase)
	if metaDatabase.ID > 0 {
		return true
	}
	return false
}

func ExistMetaDatabaseById(id int) bool {
	var metaDatabase MetaDatabase
	util.DB.Select("id").Where("id = ?", id).First(&metaDatabase)
	if metaDatabase.ID > 0 {
		return true
	}
	return false 
}

func AddMetaDatabase(data map[string]interface{}) bool {
	metaDatabase := &MetaDatabase{}
	error := util.FillStruct(data, metaDatabase)
	if error != nil {
		log.Printf("Fill Struct is Fail")
	}
	util.DB.Create(metaDatabase)

	return true
}

func EditMetaDatabase(id int, data interface{}) bool {
	util.DB.Model(&MetaDatabase{}).Where("id = ?", id).Update(data)

	return true
}

func DeleteMetaDatabase(id int) bool {
	util.DB.Where("id = ?", id).Delete(&MetaDatabase{})

	return true
}