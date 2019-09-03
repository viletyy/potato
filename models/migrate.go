package models

import (
	"github.com/viletyy/potato/models/basic"
	"github.com/viletyy/potato/pkg/util"
)

func init()  {
	util.DB.AutoMigrate(&User{}, &basic.MetaDatabase{}, &basic.Vendor{}, &basic.Business{})

	var count int
	if err := util.DB.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		//新增
		util.DB.Create(&User{
			Username: "admin",
			Password: GetSecretPassword("123456"),
			Nickname: "管理员",
		})
	}
}