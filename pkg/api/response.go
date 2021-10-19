package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回自定义状态码
const (
	Successful = 0
	Failed     = 500
)

// Response HTTP返回数据结构体, 可使用这个, 也可以自定义
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 获得一个基本的表示请求成功的 Response 对象
func Success(data interface{}) *Response {
	return &Response{Code: Successful, Data: data, Message: "success"}
}

// Error 获得一个基本的表示请求失败的 Response 对象
func Error(data interface{}) *Response {
	return &Response{Code: Failed, Data: data, Message: "fail"}
}

// WithCode 获得一个 Response 对象
func WithCode(code int, data interface{}) *Response {
	return &Response{Code: code, Data: data, Message: "success"}
}

// Msg msg 描述
func (rsp *Response) Msg(msg string) *Response {
	rsp.Message = msg
	return rsp
}

// End 结束调用
func (rsp *Response) End(c *gin.Context, httpStatus ...int) {
	status := http.StatusOK
	if len(httpStatus) > 0 {
		status = httpStatus[0]
	}

	c.JSON(status, rsp)
	return
}

// Object 直接获得本对象
func (rsp *Response) Object(_ *gin.Context) *Response {
	return rsp
}

// NewResponse 接口返回统一使用这个
//  code 服务端与客户端和web端约定的自定义状态码
//  data 具体的返回数据
//  message 可不传,自定义消息
func NewResponse(code int, data interface{}, message ...string) *Response {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	return &Response{Code: code, Data: data, Message: msg}
}
