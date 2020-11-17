package model

import (
	"go-wss/src/model/player"
)

// 玩家的数据
type Player struct {
	Model
	// 这里id本来想统一拉出来的，当时使用嵌套的Id，好像序列化有点问题...
	//Id             bson.ObjectId               `json:"_id" bson:"_id"`
	Basic    Basic                  `json:"basic"`
	ServerId int                    `json:"server_id"`
	Heroes   map[string]player.Hero `json:"Heroes"`             // 战舰
	Items    map[string]uint32      `json:"items" bson:"items"` // 道具

	Created    uint16  `json:"created"` // 创建时间
	IsWorker   bool    `json:"is_worker" bson:"is_worker"`
	TotalMoney float32 `json:"total_money" bson:"total_money"`
	Online     bool    `json:"online"` // 是否在线
}

type Basic struct {
	Lv       int    `json:"lv"`
	Exp      int    `json:"exp"`
	Name     string `json:"name"`
	ServerId int    `json:"server_id" bson:"server_id"`
}

type Energy struct {
	Pve        int16 `json:"pve" bson:"pve"`
	PveUpdated int   `json:"pve_updated" bson:"pve_updated"`

	Pvp        int16 `json:"pvp" bson:"pvp"`
	PvpUpdated int   `json:"pvp_updated" bson:"pvp_updated"`
}

func Default(serverId int) *Player {
	p := &Player{
		Basic: Basic{
			Lv:   1,
			Exp:  0,
			Name: "test",
		},
		ServerId: serverId,
	}
	return p
}
