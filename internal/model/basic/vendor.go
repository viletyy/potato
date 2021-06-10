/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 23:20:22
 * @FilePath: /potato/internal/model/basic/vendor.go
 */
package basic

import (
	"github.com/jinzhu/gorm"
	"github.com/viletyy/potato/internal/model"
)

type Vendor struct {
	model.Model

	Name string `json:"name"`
	Uuid int    `json:"uuid"`
}

func (v Vendor) Count(db *gorm.DB) (int, error) {
	var count int
	if v.Name != "" {
		db = db.Where("name = ?", v.Name)
	}
	if v.Uuid != 0 {
		db = db.Where("uuid = ?", v.Uuid)
	}
	if err := db.Model(&v).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (v Vendor) List(db *gorm.DB, pageOffset, pageSize int) (vendors []Vendor, err error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if v.Name != "" {
		db = db.Where("name = ?", v.Name)
	}
	if v.Uuid != 0 {
		db = db.Where("uuid = ?", v.Uuid)
	}
	if err = db.Find(&vendors).Error; err != nil {
		return nil, err
	}

	return
}

func (v Vendor) Get(db *gorm.DB) (vendor Vendor, err error) {
	err = db.Where("id = ?", v.ID).First(&vendor).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	return vendor, nil
}

func (v Vendor) Create(db *gorm.DB) error {
	return db.Create(&v).Error
}

func (v Vendor) Update(db *gorm.DB) error {
	return db.Model(&Vendor{}).Where("id = ?", v.ID).Update(v).Error
}

func (v Vendor) Delete(db *gorm.DB) error {
	return db.Where("id = ?", v.ID).Delete(&v).Error
}
