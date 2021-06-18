/*
 * @Date: 2021-06-18 21:37:14
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-18 22:21:54
 * @FilePath: /potato/pkg/errcode/rpc/rpc_common_code.go
 */
package rpc

var (
	RpcSuccess          = NewRpcError(0, "成功")
	RpcFail             = NewRpcError(10000, "内部错误")
	RpcInvalidParams    = NewRpcError(10001, "无效参数")
	RpcUnauthorized     = NewRpcError(10002, "认证错误")
	RpcNotFound         = NewRpcError(10003, "没有找到")
	RpcUnknown          = NewRpcError(10004, "未知")
	RpcDeadlineExceeded = NewRpcError(10005, "超出最后截止期限")
	RpcAccessDenied     = NewRpcError(10006, "访问被拒绝")
	RpcLimitExceed      = NewRpcError(10007, "访问限制")
	RpcMethodNotAllowed = NewRpcError(10008, "不支持该方法")
)
