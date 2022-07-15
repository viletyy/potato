/*
 * @Date: 2021-03-21 19:54:57
 * @LastEditors: viletyy
 * @LastEditTime: 2021-07-09 14:37:35
 * @FilePath: /potato/main.go
 */
package main

import (
	"flag"

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/initialize"
	"github.com/viletyy/potato/migrations"
	"github.com/viletyy/yolk/convert"
)

// @title Potato Api
// @version 1.0
// @description This is a potato use golang
// @BasePath /api
var grpcPort string
var httpPort string

func main() {
	global.GO_VP = initialize.Viper()
	global.GO_LOG = initialize.Zap()
	global.GO_DB = initialize.Gorm()
	global.GO_REDIS = initialize.Redis()
	global.GO_TRACER = initialize.Tracer()
	go initialize.Cron()

	if err := migrations.Migrate(global.GO_DB); err != nil {
		global.GO_LOG.Sugar().Fatalf("migrations.Migrate: %v", err)
	}

	sqlDB, err := global.GO_DB.DB()
	if err != nil {
		global.GO_LOG.Sugar().Fatalf("global.GO_DB.DB err: %v", err)
	}
	defer sqlDB.Close()
	defer global.GO_REDIS.Close()

	flag.StringVar(&grpcPort, "grpc_port", convert.ToString(global.GO_CONFIG.Server.GrpcPort), "启动grpc服务端口号")
	flag.StringVar(&httpPort, "http_port", convert.ToString(global.GO_CONFIG.Server.HttpPort), "启动http服务端口号")

	flag.Parse()

	// defaultMailer := email.NewEmail(&email.SMTPInfo{
	// 	Host:     global.GO_CONFIG.Email.Host,
	// 	Port:     global.GO_CONFIG.Email.Port,
	// 	IsSSL:    global.GO_CONFIG.Email.IsSSL,
	// 	UserName: global.GO_CONFIG.Email.UserName,
	// 	Password: global.GO_CONFIG.Email.Password,
	// 	From:     global.GO_CONFIG.Email.From,
	// })

	// _ = defaultMailer.SendMail(
	// 	global.GO_CONFIG.Email.To,
	// 	fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
	// 	fmt.Sprintf("错误信息：heheh%s", "dfds"),
	// )

	errs := make(chan error)
	go func() {
		err := initialize.RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err := initialize.RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()
	select {
	case err := <-errs:
		global.GO_LOG.Sugar().Fatalf("Run Server err: %v", err)
	}
}
