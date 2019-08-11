package basic

import "github.com/viletyy/potato/pkg/util"

func init()  {
	util.DB.AutoMigrate(&MetaDatabase{}, &Vendor{}, &Business{})
}