package models

import (
	"github.com/viletyy/potato/models/basic"
	"github.com/viletyy/potato/pkg/util"
)

func init()  {
	util.DB.AutoMigrate(&User{}, &basic.MetaDatabase{}, &basic.Vendor{}, &basic.Business{})
}