/*
 * @Date: 2021-07-09 14:33:26
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:52:23
 * @FilePath: /potato/server/user.go
 */

package server

import (
	"context"
	"encoding/json"

	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode/rpc"
	pb "github.com/viletyy/potato/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (t *UserServer) GetUserList(ctx context.Context, r *pb.GetUserListRequest) (*pb.GetUserListReply, error) {
	svc := service.New(ctx)
	dbUserList, err := svc.Dao.GetUserList(r.GetUsername(), r.GetNickname(), int(r.GetPage()), int(r.GetPageSize()))
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcErrorGetVendorListFail)
	}
	total, err := svc.Dao.CountUser(r.GetUsername(), r.GetNickname())
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcErrorCountVendorFail)
	}
	data := map[string]interface{}{
		"list": dbUserList,
		"pager": app.Pager{
			Page:     int(r.GetPage()),
			PageSize: int(r.GetPageSize()),
			Total:    total,
		},
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcInvalidParams)
	}
	userList := pb.GetUserListReply{}
	err = json.Unmarshal(byteData, &userList)
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcInvalidParams)
	}

	return &userList, nil
}
