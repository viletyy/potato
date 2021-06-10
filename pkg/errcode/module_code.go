/*
 * @Date: 2021-06-10 23:09:09
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 23:26:27
 * @FilePath: /potato/pkg/errcode/module_code.go
 */
package errcode

var (
	ErrorGetVendorListFail = NewError(21001, "获取系统厂商列表失败")
	ErrorGetVendorFail     = NewError(21002, "获取系统厂商失败")
	ErrorCreateVendorFail  = NewError(21003, "创建系统厂商失败")
	ErrorUpdateVendorFail  = NewError(21004, "更新系统厂商失败")
	ErrorDeleteVendorFail  = NewError(21005, "删除系统厂商失败")
	ErrorCountVendorFail   = NewError(21006, "统计系统厂商失败")
)
