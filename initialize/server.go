/*
 * @Date: 2021-03-22 17:03:27
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-22 23:55:32
 * @FilePath: /potato/initialize/server.go
 */
package initialize

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/routers"
	"github.com/viletyy/potato/utils"
)

var serverConfig = global.GO_CONFIG.Server

func RunServer() {
	binding.Validator = new(utils.DefaultValidator)

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverConfig.HttpPort),
		Handler:        router,
		ReadTimeout:    time.Duration(serverConfig.ReadTimeout),
		WriteTimeout:   time.Duration(serverConfig.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			global.GO_LOG.Info(fmt.Sprintf("Listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	global.GO_LOG.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		global.GO_LOG.Fatal(fmt.Sprintf("Server Shutdown: %v", err))
	}

	global.GO_LOG.Info("Server exiting")
}
