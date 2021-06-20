/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-20 19:45:38
 * @FilePath: /potato/main.go
 */
package main

import (
	"flag"
	"net"

	"github.com/soheilhy/cmux"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/initialize"
	"github.com/viletyy/yolk/convert"
)

// @title Potato Api
// @version 1.0
// @description This is a potato use golang
// @BasePath /api

var port string

func runTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func main() {
	global.GO_VP = initialize.Viper()
	global.GO_LOG = initialize.Zap()
	global.GO_DB = initialize.Gorm()
	global.GO_REDIS = initialize.Redis()
	global.GO_TRACER = initialize.Tracer()

	defer global.GO_DB.Close()
	defer global.GO_REDIS.Close()

	flag.StringVar(&port, "port", convert.ToString(global.GO_CONFIG.Server.Port), "启动端口号")
	flag.Parse()

	l, err := runTcpServer(port)
	if err != nil {
		global.GO_LOG.Sugar().Fatalf("Run Tcp Server err: %v", err)
	}
	m := cmux.New(l)
	grpcL := m.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings(
			"content-type",
			"application/grpc",
		),
	)
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := initialize.RunGrpcServer()
	httpS := initialize.RunServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		global.GO_LOG.Sugar().Fatalf("Run Server err: %v", err)
	}

}
