package main

import (
	"go-wss/src/controller"
	"go-wss/src/server"
	"log"
)

func main() {
	// 初始化路由
	controller.InitRouter()
	ws := server.NewWsServer("0.0.0.0:8000")
	log.Println(ws)
	err := ws.Start()
	if err != nil {
		log.Fatalf("启动失败: %v", err)
	}
}
