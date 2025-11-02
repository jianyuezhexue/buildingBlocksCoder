package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// HTTP工具包
type Http struct {
	ContentType string
	Headers     map[string]string
	Cookies     []*http.Cookie
	Timeout     time.Duration
}

// Options 选项函数类型
type Options func(*Http)

// 设置请求头
func WithHeaders(headers map[string]string) Options {
	return func(h *Http) {
		h.Headers = headers
	}
}

// 设置超时时间
func WithTimeout(timeout time.Duration) Options {
	return func(h *Http) {
		h.Timeout = timeout
	}
}

// 设置Cookies
func WithCookies(cookies []*http.Cookie) Options {
	return func(h *Http) {
		h.Cookies = cookies
	}
}

// 实例化HTTP工具包
func NewHttp() *Http {
	return &Http{
		ContentType: "application/json",
		Timeout:     30 * time.Second,
		Headers:     make(map[string]string),
		Cookies:     make([]*http.Cookie, 0),
	}
}

// 封装一个POST请求方法
func (h *Http) Post(url string, data any) ([]byte, error) {
	var reader io.Reader

	// 根据data类型转换为io.Reader
	switch v := data.(type) {
	case string:
		reader = strings.NewReader(v)
	case []byte:
		reader = bytes.NewReader(v)
	case io.Reader:
		reader = v
	default:
		// 对于其他类型，尝试转换为JSON
		jsonData, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data to JSON: %w", err)
		}
		reader = bytes.NewReader(jsonData)
	}

	res, err := http.Post(url, h.ContentType, reader)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
