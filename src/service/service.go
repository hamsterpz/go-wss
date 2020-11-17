package service

import (
	"fmt"
	"go-wss/src/server/client"
)

type Base struct {
	PlayerData interface{}
	Extra      interface{}
}

func (b *Base) Update() {
	fmt.Println("service.update!")
}

func (b *Base) Merge() {
	fmt.Println("service.merge!")
}

func (b *Base) MakeResponse(code uint16) *client.Response {
	return &client.Response{
		Code:  code,
		Data:  b.PlayerData,
		Extra: b.Extra,
	}
}
