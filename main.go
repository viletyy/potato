package main

import (
	"./pkg/setting"
	"./routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Potato Api
// @version 1.0
// @description This is a data_govern use golang
// @BasePath /api

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           router,
		ReadTimeout:       setting.ReadTimeout,
		WriteTimeout:      setting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf(fmt.Sprintf("Listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Printf("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Printf("Server exiting")
}