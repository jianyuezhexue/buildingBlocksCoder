package llm

type llm struct {
	AiModel      string `json:"aiModel"`      // 模型
	Url          string `json:"url"`          // 大模型接口地址
	ResponseJson bool   `json:"responseJson"` // 是否返回json
}

// options
type Option func(*llm)

// witdhResponseJson
func WithResponseJson() Option {
	return func(a *llm) {
		a.ResponseJson = true
	}
}

// 实例化AI模型
// todo 这里后期做成工厂模式，根据不同模型实例化不同的AI模型
func Newllm(opt ...Option) LLMInterface {
	model := &llm{
		Url: "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions",
	}

	for _, fn := range opt {
		fn(model)
	}
	return model
}
