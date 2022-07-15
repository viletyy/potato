/*
 * @Date: 2021-06-10 22:53:09
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-14 23:19:25
 * @FilePath: /potato/internal/dao/vendor.go
 */
package dao

import (
	"github.com/viletyy/potato/internal/model"
	"github.com/viletyy/potato/internal/model/basic"
	"github.com/viletyy/potato/pkg/app"
)

func (d *Dao) CountVendor(name string, uuid int) (int64, error) {
	vendor := basic.Vendor{Name: name, Uuid: uuid}
	return vendor.Count(d.Engine)
}

func (d *Dao) GetVendorList(name string, uuid int, page, pageSize int) ([]basic.Vendor, error) {
	vendor := basic.Vendor{Name: name, Uuid: uuid}
	pageOffset := app.GetPageOffset(page, pageSize)
	return vendor.List(d.Engine, pageOffset, pageSize)
}

func (d *Dao) GetVendor(id int64) (basic.Vendor, error) {
	vendor := basic.Vendor{
		Model: &model.Model{ID: id},
	}

	return vendor.Get(d.Engine)
}

func (d *Dao) CreateVendor(name string, uuid int) (basic.Vendor, error) {
	vendor := basic.Vendor{
		Name: name,
		Uuid: uuid,
	}

	return vendor, vendor.Create(d.Engine)
}

func (d *Dao) UpdateVendor(id int64, name string, uuid int) (basic.Vendor, error) {
	vendor := basic.Vendor{
		Name:  name,
		Uuid:  uuid,
		Model: &model.Model{ID: id},
	}

	dbVendor, err := vendor.Get(d.Engine)

	if err != nil {
		return vendor, err
	}

	return dbVendor, dbVendor.Update(d.Engine)
}

func (d *Dao) DeleteVendor(id int64) (basic.Vendor, error) {
	vendor := basic.Vendor{
		Model: &model.Model{ID: id},
	}

	dbVendor, err := vendor.Get(d.Engine)

	if err != nil {
		return vendor, err
	}

	return dbVendor, dbVendor.Delete(d.Engine)
}
