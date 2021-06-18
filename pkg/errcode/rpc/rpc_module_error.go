/*
 * @Date: 2021-06-18 22:32:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-18 22:33:40
 * @FilePath: /potato/pkg/errcode/rpc/rpc_module_error.go
 */
package rpc

var (
	RpcErrorGetVendorListFail = NewRpcError(20101, "获取系统厂商列表失败")
	RpcErrorGetVendorFail     = NewRpcError(20102, "获取系统厂商失败")
	RpcErrorCreateVendorFail  = NewRpcError(20103, "创建系统厂商失败")
	RpcErrorUpdateVendorFail  = NewRpcError(20104, "更新系统厂商失败")
	RpcErrorDeleteVendorFail  = NewRpcError(20105, "删除系统厂商失败")
	RpcErrorCountVendorFail   = NewRpcError(20106, "统计系统厂商失败")
)
