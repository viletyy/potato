/*
 * @Date: 2021-06-17 00:19:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:24:14
 * @FilePath: /potato/initialize/grpc_server.go
 */
package initialize

import (
	"net"

	pb "github.com/viletyy/potato/proto/basic"
	"github.com/viletyy/potato/server/basic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(port string) error {
	server := grpc.NewServer()
	pb.RegisterVendorServiceServer(server, basic.NewVendorServer())
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return server.Serve(lis)
}
