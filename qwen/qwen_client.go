package qwen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xinggaoya/qwen-sdk/qwenmodel"
	"net/http"
)

type Chat struct {
	BaseUrl   string
	ApiKey    string
	QWenModel string
	Params    qwenmodel.Parameters
}

func NewWithDefaultChat(apiKey string) *Chat {
	return &Chat{
		BaseUrl:   qwenmodel.ChatBaseUrl,
		ApiKey:    apiKey,
		QWenModel: qwenmodel.ChatQWenModel,
		Params:    qwenmodel.Parameters{EnableSearch: true},
	}
}

// GetAIReply 获取聊天回复
func (c *Chat) GetAIReply(messages []qwenmodel.Messages) qwenmodel.Response {
	client := http.Client{}

	if !checkParams(c) {
		return qwenmodel.Response{}
	}
	// body
	body := qwenmodel.QWenTurbo{
		Model:      c.QWenModel,
		Input:      qwenmodel.Input{Messages: messages},
		Parameters: c.Params,
	}
	jsonBody, err := json.Marshal(body)
	// 创建请求
	req, err := http.NewRequest("POST", c.BaseUrl, bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Printf("NewRequest failed,err:%v\n", err)
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	// 发送请求
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("client.Do failed,err:%v\n", err)
	}
	defer resp.Body.Close()

	// 读取响应
	var result qwenmodel.Response

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		fmt.Printf("json.NewDecoder failed,err:%v\n", err)
	}

	return result
}

// 效验参数
func checkParams(chat *Chat) bool {
	if chat.QWenModel == "" {
		fmt.Errorf("QWenModel is empty")
		return false
	}
	if chat.ApiKey == "" {
		fmt.Errorf("ApiKey is empty")
		return false
	}
	if chat.BaseUrl == "" {
		fmt.Errorf("BaseUrl is empty")
		return false
	}
	return true
}
