package player

import (
	"fmt"
	"go-wss/src/config"
	"go-wss/src/model"
	"go-wss/src/server/client"
	"go-wss/src/service"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func Login(params client.Params) *client.Response {
	// 获取玩家的参数
	userId := params.Data["uid"].(string)
	// 登录
	fmt.Println(config.System)
	p := model.Player{}
	_ = model.CPlayer.FindId(bson.ObjectIdHex(userId)).One(&p)
	// 服务类
	d := service.Base{PlayerData: p}
	log.Println("有玩家登陆啦！", p)
	params.Client.SetUid(userId)
	params.Client.SetServerId(p.ServerId)
	client.Members.Add(params.Client)

	return d.MakeResponse(200)
}

func ActionChat(params client.Params) *client.Response {
	serverId, _ := params.Data["server_id"].(int)
	uid, _ := params.Data["uid"].(string)

	targetClient, ok := client.Members.TargetClient(serverId, uid)
	if ok {
		targetClient.MsgChan <- &client.Response{Data: "给你发的消息"}
	}
	return nil
}
