package main

import (
	"context"
	_ "eassy/core/config"
	_ "eassy/core/etcd"
	"eassy/model"
	"eassy/server"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

var (
	serverType string
	serverPort string
	help       bool
)

func main() {
	//flag 命令解析
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	//go getGoroutineNum()
	model.InitDB(serverType)
	s := server.NewServer(serverType, serverPort)
	go func() {
		if err := server.Run(s); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Start Error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
	log.Println("Server Shutdown")
}

func init() {
	flag.StringVar(&serverType, "t", "login", "server type")
	flag.StringVar(&serverPort, "p", "5566", "server port")
	flag.BoolVar(&help, "h", false, "帮助信息")
}

func getGoroutineNum() {
	var t = time.NewTicker(1000 * time.Millisecond)
	for {
		select {
		case <-t.C:
			fmt.Println("============当前协程数：", runtime.NumGoroutine(), "=============")
		}
	}
}
