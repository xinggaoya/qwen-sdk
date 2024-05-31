package qwen

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Chat struct {
	BaseUrl   string
	ApiKey    string
	QWenModel string
	Params    Parameters
}

func NewWithDefaultChat(apiKey string) *Chat {
	return &Chat{
		BaseUrl:   ModelBaseUrl,
		ApiKey:    apiKey,
		QWenModel: ModelQWenTurbo,
		Params:    Parameters{EnableSearch: true, ResponseFormat: "message"},
	}
}

// GetAIReply 获取聊天回复
func (c *Chat) GetAIReply(messages []Messages) (Response, error) {
	client := http.Client{}

	if !checkParams(c) {
		return Response{}, errors.New("invalid parameters")
	}
	// body
	body := QWenTurbo{
		Model:      c.QWenModel,
		Input:      Input{Messages: messages},
		Parameters: c.Params,
	}
	jsonBody, err := json.Marshal(body)
	// 创建请求
	req, err := http.NewRequest("POST", c.BaseUrl, bytes.NewReader(jsonBody))
	if err != nil {
		return Response{}, fmt.Errorf("http.NewRequest failed,err:%v", err)
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	// 发送请求
	resp, err := client.Do(req)

	if err != nil {
		return Response{}, fmt.Errorf("http.Client.Do failed,err:%v", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// 读取响应
	var result Response

	if resp.StatusCode != http.StatusOK {
		var errResp ResponseError
		err = json.Unmarshal(b, &errResp)
		if err != nil {
			return Response{}, err
		}
		return Response{}, fmt.Errorf("failed,err:%v,code:%s,message:%s", err, errResp.Code, errResp.Message)
	}

	err = json.Unmarshal(b, &result)

	if err != nil {
		fmt.Printf("json.NewDecoder failed,err:%v\n", err)
	}

	return result, nil
}

// GetAIReplyStream 获取聊天回复
func (c *Chat) GetAIReplyStream(messages []Messages) (<-chan Response, error) {
	client := http.Client{}

	if !checkParams(c) {
		return nil, fmt.Errorf("invalid parameters")
	}

	// Prepare request body
	body := QWenTurbo{
		Model:      c.QWenModel,
		Input:      Input{Messages: messages},
		Parameters: c.Params,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", c.BaseUrl, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("new request failed: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-DashScope-SSE", "enable")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ResponseError
		b, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(b, &errResp)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed,err:%v,code:%s,message:%s", err, errResp.Code, errResp.Message)
	}

	// Handle streaming response
	messageChan := make(chan Response)
	go func() {
		info := ""
		defer resp.Body.Close()
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(messageChan)
					return
				}
				fmt.Fprintf(os.Stderr, "Error reading stream: %v\n", err)
				close(messageChan)
				return
			}
			// Remove trailing newline if present
			if line[len(line)-1] == '\n' {
				line = line[:len(line)-1]
			}
			// 只获取前锥 data: 前缀
			if len(line) > 5 && line[:5] == "data:" {
				// 解析json
				var result Response
				err = json.Unmarshal([]byte(line[5:]), &result)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
				}
				if info != result.Output.Text {
					messageChan <- result
				}
			}
		}
	}()

	return messageChan, nil
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
