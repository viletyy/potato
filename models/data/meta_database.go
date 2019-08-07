package data

import "github.com/viletyy/potato/pkg/util"

type MetaDatabase struct {
	util.Model

	Name string `json:"name"`
	CnName string `json:"cn_name"`
	Logo string `json:"logo"`
}

func GetMetaDatabases(pageNum int, pageSize int, maps interface{}) (metaDatabases []MetaDatabase) {
	util.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&metaDatabases)
	return
}

func GetMetaDatabaseTotal(maps interface{}) (count int) {
	util.DB.Model(&MetaDatabase{}).Where(maps).Count(&count)

	return
}

