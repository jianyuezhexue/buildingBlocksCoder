package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LLMInterface interface {
	Request(prompt, input string) (string, error)
}

// AI请求封装
func (a *llm) Request(sysPrompt, userInput string) (string, error) {

	// todo 后期根据不同模型处理不同逻辑

	// 创建 HTTP 客户端
	client := &http.Client{}
	// 构建请求体
	requestBody := RequestBody{
		Model: "qwen-plus",
		Messages: []*Message{
			{
				Role:    "system",
				Content: sysPrompt,
			},
			{
				Role:    "user",
				Content: userInput,
			},
		},
		ResponseFormat: &ResponseFormat{
			Type: "text",
		},
	}

	// 指定返回了类型
	if a.ResponseJson {
		requestBody.ResponseFormat = &ResponseFormat{
			Type: "json_object",
		}
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", nil
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", a.Url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", nil
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer sk-e692504205e74522b45710e1c25065ad")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	res := DeepSeekResp{}
	// 解析响应内容
	if err := json.Unmarshal(bodyText, &res); err != nil {
		return "", nil
	}

	// 异常拦截
	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}

	// todo 后期这里将提示词，返回结果等做成日志存储

	return res.Choices[0].Message.Content, nil
}
