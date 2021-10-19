package http

import (
	"bytes"
	"encoding/json"
	"go-project/pkg/logger"
	"io"
	"net/http"
	"time"
)

// HttpGet 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func HttpGet(url string) string {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if resp == nil {
		logger.Logger("http.get").Error(url + "接口超时！")
		return ""
	}
	if err != nil {
		logger.Logger("http.get").Error(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			logger.Logger("http.get").Error(err)
		}
	}
	return result.String()
}

// HttpPost 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func HttpPost(url string, data interface{}, contentType string) string {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if resp == nil {
		logger.Logger("http.post").Error(url + "接口超时！")
		return ""
	}
	if err != nil {
		logger.Logger("http.post").Error(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			logger.Logger("http.post").Error(err)
		}
	}
	return result.String()
}
