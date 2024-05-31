package main

import (
	"fmt"
	"github.com/xinggaoya/qwen-sdk/qwen"
	"testing"
)

func TestRun(t *testing.T) {

	// 初始化QWEN聊天机器人客户端，使用您的API密钥
	apiKey := "your api key"
	qwenclient := qwen.NewWithDefaultChat(apiKey)

	//qwenclient.QWenModel = qwen.ModelQWenMax

	// 定义一条消息对话的历史记录
	messages := []qwen.Messages{
		{Role: qwen.ChatUser, Content: "你好"},
		{Role: qwen.ChatBot, Content: "你好！有什么我能为你做的吗？"},
		{Role: qwen.ChatUser, Content: "我想买一件衬衫"},
	}

	// 获取AI对消息的回复
	resp, err := qwenclient.GetAIReply(messages)
	if err != nil {
		fmt.Printf("获取AI回复失败：%v\n", err)
		return
	}

	// 打印收到的回复
	fmt.Printf("收到的回复：%v\n", resp.Output.Text)

}

func TestRunStream(t *testing.T) {
	apiKey := "your api key"
	qwenclient := qwen.NewWithDefaultChat(apiKey)

	//qwenclient.QWenModel = "new model"

	// 定义一条消息对话的历史记录
	messages := []qwen.Messages{
		{Role: qwen.ChatUser, Content: "你好"},
		{Role: qwen.ChatBot, Content: "你好！有什么我能为你做的吗？"},
		{Role: qwen.ChatUser, Content: "我想买一件衬衫"},
	}

	// 获取AI对消息的回复
	resp, err := qwenclient.GetAIReplyStream(messages)

	if err != nil {
		fmt.Printf("获取AI回复失败：%v\n", err)
		return
	}
	// 打印收到的回复
	for info := range resp {
		fmt.Printf("收到的回复：%v\n", info.Output.Text)
	}
}
