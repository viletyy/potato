/*
 * @Date: 2021-03-22 17:03:27
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 22:24:39
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

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/routers"
)

func RunServer() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.GO_CONFIG.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    time.Duration(global.GO_CONFIG.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.GO_CONFIG.Server.ReadTimeout) * time.Second,
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
