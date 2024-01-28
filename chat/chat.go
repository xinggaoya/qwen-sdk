package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xinggaoya/qwen-sdk/constant"
	"github.com/xinggaoya/qwen-sdk/model"
	"net/http"
)

type Chat struct {
	BaseUrl string
	ApiKey  string
	Model   string
}

func NewDefaultChat(apiKey string) *Chat {
	return &Chat{
		BaseUrl: apiKey,
		ApiKey:  constant.ChatApiKey,
		Model:   constant.ChatModel,
	}
}

// GetChatAnswer 获取聊天回复
func (c *Chat) GetChatAnswer(messages []model.Messages) model.Response {
	client := http.Client{}

	// body
	body := model.QWenTurbo{
		Model: constant.ChatModel,
		Input: model.Input{Messages: messages},
		Parameters: model.Parameters{
			EnableSearch: true,
		},
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
	var result model.Response

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		fmt.Printf("json.NewDecoder failed,err:%v\n", err)
	}

	return result
}
