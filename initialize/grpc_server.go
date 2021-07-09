/*
 * @Date: 2021-06-17 00:19:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:52:03
 * @FilePath: /potato/initialize/grpc_server.go
 */
package initialize

import (
	"net"

	pb "github.com/viletyy/potato/proto"
	basic_pb "github.com/viletyy/potato/proto/basic"
	grpc_server "github.com/viletyy/potato/server"
	"github.com/viletyy/potato/server/basic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(port string) error {
	server := grpc.NewServer()
	basic_pb.RegisterVendorServiceServer(server, basic.NewVendorServer())
	pb.RegisterUserServiceServer(server, grpc_server.NewUserServer())

	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return server.Serve(lis)
}
