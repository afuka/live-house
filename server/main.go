package main

import (
	"afkser/facades"
	"afkser/initialize"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 初始化其他资源具柄
	initialize.Init()

	// 门面
	facades.DB = initialize.DB
	facades.Redis = initialize.Redis

	// gin 重启或停止你的Web服务器 等待当前所有的都完毕后再重启或停止
	r := initialize.NewRouter()

	srv := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
