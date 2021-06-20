/*
 * @Date: 2021-03-22 17:03:27
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-20 22:22:10
 * @FilePath: /potato/initialize/http_server.go
 */
package initialize

import (
	"net/http"
	"time"

	"github.com/viletyy/potato/global"
	"github.com/viletyy/potato/internal/routers"
)

func RunHttpServer(port string) *http.Server {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    time.Duration(global.GO_CONFIG.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.GO_CONFIG.Server.ReadTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}
