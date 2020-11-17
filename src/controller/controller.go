package controller

import (
	"go-wss/src/controller/player"
	"go-wss/src/server/client"
)

const (
	ActionLogin = client.ActionLogin // 登录
	ActionChat  = "chat"
)

func InitRouter() {
	client.ActionMap[ActionLogin] = player.Login
	client.ActionMap[ActionChat] = player.ActionChat
}
