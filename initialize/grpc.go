/*
 * @Date: 2021-06-17 00:19:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-17 00:33:38
 * @FilePath: /potato/initialize/grpc.go
 */
package initialize

import (
	"net"

	"github.com/viletyy/potato/global"
	pb "github.com/viletyy/potato/proto/basic"
	"github.com/viletyy/potato/server/basic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpc() {
	server := grpc.NewServer()
	pb.RegisterVendorServiceServer(server, basic.NewVendorServer())
	reflection.Register(server)

	go func() {
		listen, err := net.Listen("tcp", ":10002")
		if err != nil {
			global.GO_LOG.Sugar().Fatalf("net.Listen err: %v", err)
		}

		err = server.Serve(listen)
		if err != nil {
			global.GO_LOG.Sugar().Fatalf("server.Serve err: %v", err)
		}
	}()
}
