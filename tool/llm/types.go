package llm

// AI如惨请求
type AiInput struct {
	Input string `json:"input"`
}

// AI请求
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

// 千问请求
type RequestBody struct {
	Model          string          `json:"model"`
	Messages       []*Message      `json:"messages"`
	ResponseFormat *ResponseFormat `json:"response_format"`
}

// 千问返回
type DeepSeekResp struct {
	Choices           []*Choice `json:"choices"`
	Object            string    `json:"object"`
	Usage             Usage     `json:"usage"`
	Created           int64     `json:"created"`
	SystemFingerprint *string   `json:"system_fingerprint,omitempty"`
	Model             string    `json:"model"`
	ID                string    `json:"id"`
	Error             *Error    `json:"error,omitempty"`
}
type Error struct {
	Code    string `json:"code"`
	Param   string `json:"param"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Choice struct {
	Message      ResMessage   `json:"message"`
	FinishReason string       `json:"finish_reason"`
	Index        int          `json:"index"`
	Logprobs     *interface{} `json:"logprobs,omitempty"` // 使用 interface{} 因为 logprobs 的值是 null
}

type ResMessage struct {
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content"`
	Role             string `json:"role"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
