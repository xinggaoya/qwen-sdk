package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	// 初始化QWEN聊天机器人客户端，使用您的API密钥
	apiKey := "your api key"
	qwenclient := NewWithDefaultChat(apiKey)

	//qwenclient.QWenModel = "new model"

	// 定义一条消息对话的历史记录
	messages := []Messages{
		{Role: ChatUser, Content: "你好"},
		{Role: ChatBot, Content: "你好！有什么我能为你做的吗？"},
		{Role: ChatUser, Content: "我想买一件衬衫"},
	}

	// 获取AI对消息的回复
	resp := qwenclient.GetAIReply(messages)

	// 打印收到的回复
	fmt.Printf("收到的回复：%v\n", resp.Output.Text)

}

func TestName(t *testing.T) {
	apiKey := "sk-cf5052b3c8314dfeabd520c700e55869"
	qwenclient := NewWithDefaultChat(apiKey)

	//qwenclient.QWenModel = "new model"

	// 定义一条消息对话的历史记录
	messages := []Messages{
		{Role: ChatUser, Content: "你好"},
		{Role: ChatBot, Content: "你好！有什么我能为你做的吗？"},
		{Role: ChatUser, Content: "我想买一件衬衫"},
	}

	// 获取AI对消息的回复
	resp, _ := qwenclient.GetAIReplyStream(messages)

	// 打印收到的回复
	for msg := range resp {
		fmt.Printf("收到的回复：%v\n", msg)
	}
}
