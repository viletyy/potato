/*
 * @Date: 2021-06-17 00:19:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-20 19:06:15
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

	// go func() {
	// 	listen, err := net.Listen("tcp", ":"+convert.ToString(global.GO_CONFIG.Server.RpcPort))
	// 	if err != nil {
	// 		global.GO_LOG.Sugar().Fatalf("net.Listen err: %v", err)
	// 	}

	// 	err = server.Serve(listen)
	// 	if err != nil {
	// 		global.GO_LOG.Sugar().Fatalf("server.Serve err: %v", err)
	// 	}
	// }()
	return server
}
