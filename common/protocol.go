package common

import "encoding/json"

type Job struct {
	Name    string
	Command string
	CornExp string
}

type Response struct {
	ErrorCode int         `json:"errorCode"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

// BuildResponse 构建 response 方法
func BuildResponse(errorCode int, msg string, data interface{}) ([]byte, error) {
	response := &Response{
		ErrorCode: errorCode,
		Msg:       msg,
		Data:      data,
	}
	return json.Marshal(response)
}