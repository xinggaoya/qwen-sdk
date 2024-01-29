# Qwen SDK

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/xinggaoya/qwen-sdk)](https://github.com/xinggaoya/qwen-sdk/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/xinggaoya/qwen-sdk)](https://goreportcard.com/report/github.com/xinggaoya/qwen-sdk)
[![GitHub license](https://img.shields.io/github/license/xinggaoya/qwen-sdk)](LICENSE)

Qwen SDK 是一个专为开发者打造的便捷工具包，提供了一系列功能强大的API和工具函数，帮助您轻松实现与通义千问平台的集成。

## 特性

- **简单易用**：Qwen SDK 提供了简单易用的API，帮助您轻松实现与通义千问平台的集成。
- **功能丰富**：Qwen SDK 提供了丰富的API，涵盖了通义千问平台的各项功能。
- **持续更新**：Qwen SDK 持续跟进通义千问平台的最新功能，为您提供全面的开发支持。
  
## 安装

在您的Go项目中通过 `go get` 命令安装Qwen SDK：

```sh
go get github.com/xinggaoya/qwen-sdk
```

或者将它添加到您的项目 `go.mod` 文件中的依赖列表：

```sh
require github.com/xinggaoya/qwen-sdk v1.0.0
```

## 快速开始

```go
import "github.com/xinggaoya/qwen-sdk"

func main() {
   // 初始化QWEN聊天机器人客户端，使用您的API密钥
   apiKey := "your api key"
   qwenclient := chat.NewWithDefaultChat(apiKey)

   //qwenclient.QWenModel = "new model"

   // 定义一条消息对话的历史记录
   messages := []qwenmodel.Messages{
   {Role: qwenmodel.ChatUser, Content: "你好"},
   {Role: qwenmodel.ChatBot, Content: "你好！有什么我能为你做的吗？"},
   {Role: qwenmodel.ChatUser, Content: "我想买一件衬衫"},
   }

   // 获取AI对消息的回复
   resp := qwenclient.GetAIReply(messages)

   // 打印收到的回复
   fmt.Printf("收到的回复：%v\n", resp.Output.Text)
}

```

## 文档

详细文档及API参考，请访问 [Qwen SDK Docs](https://qwen.github.io/sdk/docs)。

## 示例

查看项目中的 `/build_test` 目录获取更多使用示例。

## 贡献

欢迎贡献代码、报告问题或提出改进建议！请参阅 [CONTRIBUTING.md](CONTRIBUTING.md) 获取更多关于参与本项目的指南。

## 许可证

Qwen SDK 遵循 MIT 许可证。有关详细信息，请参阅 LICENSE 文件。

---
