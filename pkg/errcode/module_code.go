/*
 * @Date: 2021-06-10 23:09:09
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:32:17
 * @FilePath: /potato/pkg/errcode/module_code.go
 */
package errcode

var (
	ErrorUploadFileFail    = NewError(20001, "上传文件失败")
	ErrorRegisterUserError = NewError(20010, "用户注册失败")
	ErrorLoginUserError    = NewError(20011, "用户登录失败")
	ErrorGetUserListFail   = NewError(20101, "获取用户列表失败")
	ErrorGetUserFail       = NewError(20102, "获取用户失败")
	ErrorCreateUserFail    = NewError(20103, "创建用户失败")
	ErrorUpdateUserFail    = NewError(20104, "更新用户失败")
	ErrorDeleteUserFail    = NewError(20105, "删除用户失败")
	ErrorCountUserFail     = NewError(20106, "统计用户失败")
	ErrorGetVendorListFail = NewError(20201, "获取系统厂商列表失败")
	ErrorGetVendorFail     = NewError(20202, "获取系统厂商失败")
	ErrorCreateVendorFail  = NewError(20203, "创建系统厂商失败")
	ErrorUpdateVendorFail  = NewError(20204, "更新系统厂商失败")
	ErrorDeleteVendorFail  = NewError(20205, "删除系统厂商失败")
	ErrorCountVendorFail   = NewError(20206, "统计系统厂商失败")
)
