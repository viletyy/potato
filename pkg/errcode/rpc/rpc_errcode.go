/*
 * @Date: 2021-06-18 21:42:53
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-18 22:31:20
 * @FilePath: /potato/pkg/errcode/rpc/rpc_errcode.go
 */
package rpc

import (
	"fmt"

	pb "github.com/viletyy/potato/proto"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type RpcError struct {
	Code int
	Msg  string
}

type Status struct {
	*status.Status
}

var rpcCodes = map[int]string{}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func NewRpcError(code int, msg string) *RpcError {
	if _, ok := rpcCodes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	rpcCodes[code] = msg
	return &RpcError{Code: code, Msg: msg}
}

func (e *RpcError) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code, e.Msg)
}

func ToRpcError(err *RpcError) error {
	pbErr := &pb.Error{Code: int32(err.Code), Message: err.Msg}
	s, _ := status.New(ToRpcCode(err.Code), err.Msg).WithDetails(pbErr)
	return s.Err()
}

func ToRpcStatus(code int, msg string) *Status {
	pbErr := &pb.Error{Code: int32(code), Message: msg}
	s, _ := status.New(ToRpcCode(code), msg).WithDetails(pbErr)
	return &Status{s}
}

func ToRpcCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case RpcFail.Code:
		statusCode = codes.Internal
	case RpcInvalidParams.Code:
		statusCode = codes.InvalidArgument
	case RpcUnauthorized.Code:
		statusCode = codes.Unauthenticated
	case RpcAccessDenied.Code:
		statusCode = codes.PermissionDenied
	case RpcDeadlineExceeded.Code:
		statusCode = codes.DeadlineExceeded
	case RpcNotFound.Code:
		statusCode = codes.NotFound
	case RpcLimitExceed.Code:
		statusCode = codes.ResourceExhausted
	case RpcMethodNotAllowed.Code:
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}

	return statusCode
}
