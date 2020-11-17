package client

import (
	"github.com/gorilla/websocket"
	"strconv"
)

type Client struct {
	Conn     *websocket.Conn
	Uid      string
	ServerId int
	MsgChan  chan *Response
}

func (c *Client) SetUid(uid string) {
	// 设置登录的uid
	c.Uid = uid
}

func (c *Client) SetServerId(serverId int) {
	// 设置服务器id
	c.ServerId = serverId
}

type members struct {
	Clients map[string]map[string]*Client
	Count   int
}

func (m *members) Add(client *Client) {
	// 新建一个连接
	serverIdStr := strconv.Itoa(client.ServerId)
	if m.Clients == nil {
		m.Clients = make(map[string]map[string]*Client)
	}
	if _, ok := m.Clients[serverIdStr]; !ok {
		d := make(map[string]*Client)
		m.Clients[serverIdStr] = d
	}
	m.Clients[serverIdStr][client.Uid] = client
}

func (m *members) ServerClients(serverId int) map[string]*Client {
	serverIdStr := strconv.Itoa(serverId)
	if serverClients, ok := m.Clients[serverIdStr]; ok {
		return serverClients
	} else {
		return make(map[string]*Client)
	}
}

func (m *members) TargetClient(serverId int, uid string) (*Client, bool) {
	clients := m.ServerClients(serverId)
	if c, ok := clients[uid]; ok {
		return c, true
	}
	return nil, false
}

func (m *members) IsOnline(serverId int, uid string) bool {
	if _, ok := m.TargetClient(serverId, uid); ok {
		return true
	}
	return false
}

// 根据服务id来区分的 {"服务器id": {"玩家id": "连接的数据"}}
var Members = members{}
