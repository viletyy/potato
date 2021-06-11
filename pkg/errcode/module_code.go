/*
 * @Date: 2021-06-10 23:09:09
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-11 17:25:58
 * @FilePath: /potato/pkg/errcode/module_code.go
 */
package errcode

var (
	ErrorUploadFileFail    = NewError(20001, "上传文件失败")
	ErrorGetVendorListFail = NewError(20101, "获取系统厂商列表失败")
	ErrorGetVendorFail     = NewError(20102, "获取系统厂商失败")
	ErrorCreateVendorFail  = NewError(20103, "创建系统厂商失败")
	ErrorUpdateVendorFail  = NewError(20104, "更新系统厂商失败")
	ErrorDeleteVendorFail  = NewError(20105, "删除系统厂商失败")
	ErrorCountVendorFail   = NewError(20106, "统计系统厂商失败")
)
