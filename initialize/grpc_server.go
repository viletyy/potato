/*
 * @Date: 2021-06-17 00:19:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-20 19:46:17
 * @FilePath: /potato/initialize/grpc_server.go
 */
package initialize

import (
	pb "github.com/viletyy/potato/proto/basic"
	"github.com/viletyy/potato/server/basic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterVendorServiceServer(server, basic.NewVendorServer())
	reflection.Register(server)

	return server
}
