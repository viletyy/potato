package basic

import (
	"github.com/viletyy/potato/pkg/util"
)

type Business struct {
	util.Model

	Name string `json:"name"`
	Desc string `json:"desc"`
	CId int `json:"c_id"`
}

func GetBusinesses(pageNum int, pageSize int, maps interface{}) (businesses []Business) {
	util.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&businesses)
	return 
}

func GetBusinessTotal(maps interface{}) (count int) {
	util.DB.Model(&Business{}).Where(maps).Count(&count)
	return
}

func ExistBusinessByName(name string) bool {
	var business Business
	util.DB.Select("id").Where("name = ?", name).First(&business)
	if business.ID > 0 {
		return true
	}
	return false
}

func ExistBusinessById(id int) bool {
	var business Business
	util.DB.Select("id").Where("id = ?", id).First(&business)
	if business.ID > 0 {
		return true
	}
	return false
}

func AddBusiness(name string, desc string, cId int) bool {
	util.DB.Create(&Business {
		Name: name,
		Desc: desc,
		CId: cId,
	})

	return true
}

func EditBusiness(id int, data interface{}) bool {
	util.DB.Model(&Business{}).Where("id = ?", id).Update(data)

	return true
}

func DeleteBusiness(id int) bool {
	util.DB.Where("id = ?", id).Delete(&Business{})

	return true
}