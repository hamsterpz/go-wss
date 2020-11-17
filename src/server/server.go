package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-wss/src/server/client"
	"go-wss/src/server/errors"
	"net"
	"net/http"
)

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

func NewWsServer(addr string) *WsServer {
	ws := new(WsServer)
	ws.addr = addr
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}
	return ws
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		httpCode := http.StatusInternalServerError
		body := http.StatusText(httpCode)
		fmt.Println("路由错误", body)
		http.Error(w, body, httpCode)
		return
	}
	// 升级为websocket
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	fmt.Println("client connect :", conn.RemoteAddr())

	msgChan := make(chan *client.Response)
	go ws.connHandle(&client.Client{Conn: conn, MsgChan: msgChan})

}

func (ws *WsServer) connHandle(c *client.Client) {
	defer func() {
		_ = c.Conn.Close()
	}()

	stopCh := make(chan int)
	go ws.send(c.Conn, stopCh, c.MsgChan)

	for {
		//conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(5000)))
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			close(stopCh)
			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					fmt.Printf("ReadMessage timeout remote: %v\n", c.Conn.RemoteAddr())
					return
				}
			}
			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("ReadMessage other remote:%v error: %v \n", c.Conn.RemoteAddr(), err)
			}
			return
		}

		// 如果注册了函数
		params := client.Params{Client: c}
		// 如果解析的数据格式错误，则直接返回错误码
		unmarshalErr := json.Unmarshal(msg, &params)
		if unmarshalErr != nil {
			c.MsgChan <- client.NotImplement
		} else {
			if actionFunc, ok := client.ActionMap[params.Action]; ok {
				// T封装一下参数
				if params.Action == client.ActionLogin || c.Uid != "" {
					// 如果是登录，或者已经登录过了
					d, _ := ErrorInterpreter(actionFunc, params)
					c.MsgChan <- d
				} else {
					c.MsgChan <- client.NotImplement
				}
			} else {
				c.MsgChan <- client.LoginRequired
			}
		}
	}
}

func (ws *WsServer) send(conn *websocket.Conn, stopCh chan int, msgChannel chan *client.Response) {
	for {
		select {
		case <-stopCh:
			// TODO 设置下线的操作
			fmt.Println("connect closed")
			return
		case d := <-msgChannel:
			msg, _ := json.Marshal(d)
			err := conn.WriteMessage(1, msg)
			if err != nil {
				fmt.Println("send msg faild ", err)
				return
			}
		}

	}
}

func (ws *WsServer) Start() (err error) {
	ws.listener, err = net.Listen("tcp", ws.addr)
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}
	err = http.Serve(ws.listener, ws)
	if err != nil {
		fmt.Println("http serve error:", err)
		return
	}
	return nil
}

/*
全局错误的捕获(用于错误码)
*/

func ErrorInterpreter(f func(params client.Params) *client.Response, params client.Params) (res *client.Response, err errors.Error) {
	defer func() {
		if r := recover(); r != nil {
			// 创建一个新的
			res = &client.Response{}
			if r, ok := r.(errors.Error); ok {
				err = r
				res.Code = err.Code
				res.Extra = params.Data
			} else {
				err = errors.Error{Code: 500}
			}

		}
	}()
	res = f(params)
	return res, err
}
