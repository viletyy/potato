/*
 * @Date: 2021-06-10 16:25:21
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 23:52:39
 * @FilePath: /potato/pkg/errcode/common_code.go
 */
package errcode

var (
	Success                   = NewError(0, "请求成功")
	ServerError               = NewError(10000, "服务内部错误")
	InvalidParams             = NewError(10001, "参数错误")
	InvalidMultiParams        = NewError(10002, "多个参数错误")
	NotFound                  = NewError(10003, "找不到数据")
	UnauthorizedAuthNotExist  = NewError(10004, "鉴权失败，找不到对应等 AppKey和 AppSecret")
	UnauthorizedTokenError    = NewError(10005, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10006, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10007, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10008, "请求过多")
)
