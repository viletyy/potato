/*
 * @Date: 2021-06-17 00:19:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-19 22:41:24
 * @FilePath: /potato/initialize/grpc.go
 */
package initialize

import (
	"net"

	"github.com/viletyy/potato/global"
	pb "github.com/viletyy/potato/proto/basic"
	"github.com/viletyy/potato/server/basic"
	"github.com/viletyy/yolk/convert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpc() {
	server := grpc.NewServer()
	pb.RegisterVendorServiceServer(server, basic.NewVendorServer())
	reflection.Register(server)

	go func() {
		listen, err := net.Listen("tcp", ":"+convert.ToString(global.GO_CONFIG.Server.RpcPort))
		if err != nil {
			global.GO_LOG.Sugar().Fatalf("net.Listen err: %v", err)
		}

		err = server.Serve(listen)
		if err != nil {
			global.GO_LOG.Sugar().Fatalf("server.Serve err: %v", err)
		}
	}()
}
