package client

import "go-wss/src/server/errors"

// 请求的参数
type Params struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
	Client *Client
}

// 响应的数据
type Response struct {
	Code  uint16      `json:"code"`
	Data  interface{} `json:"data"`
	Extra interface{} `json:"extra"`
}

// 未实现接口或者参数错误
var NotImplement = &Response{Code: errors.NotImplement}
var LoginRequired = &Response{Code: errors.LoginRequired}

// 路由注册
var ActionMap = map[string]func(params Params) *Response{}

// 登录接口
const ActionLogin = "login"
