package basic

import (
	"github.com/viletyy/potato/pkg/util"
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
