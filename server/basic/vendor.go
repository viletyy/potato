/*
 * @Date: 2021-06-16 23:49:02
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-18 22:39:40
 * @FilePath: /potato/server/basic/vendor.go
 */
package basic

import (
	"context"
	"encoding/json"

	"github.com/viletyy/potato/internal/service"
	"github.com/viletyy/potato/pkg/app"
	"github.com/viletyy/potato/pkg/errcode/rpc"
	pb "github.com/viletyy/potato/proto/basic"
)

type VendorServer struct {
	pb.UnimplementedVendorServiceServer
}

func NewVendorServer() *VendorServer {
	return &VendorServer{}
}

func (t *VendorServer) GetVendorList(ctx context.Context, r *pb.GetVendorListRequest) (*pb.GetVendorListReply, error) {
	svc := service.New(ctx)
	dbVendorList, err := svc.Dao.GetVendorList(r.GetName(), int(r.GetUuid()), int(r.GetPage()), int(r.GetPageSize()))
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcErrorGetVendorListFail)
	}
	total, err := svc.Dao.CountVendor(r.GetName(), int(r.GetUuid()))
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcErrorCountVendorFail)
	}
	data := map[string]interface{}{
		"list": dbVendorList,
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
	vendorList := pb.GetVendorListReply{}
	err = json.Unmarshal(byteData, &vendorList)
	if err != nil {
		return nil, rpc.ToRpcError(rpc.RpcInvalidParams)
	}
	return &vendorList, nil
}
