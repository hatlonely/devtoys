package devtoys

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpClientApp struct {
	ctx context.Context
}

func NewHttpClientApp() *HttpClientApp {
	return &HttpClientApp{}
}

func (a *HttpClientApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type HttpClientReq struct {
	Method        string
	URL           string
	Query         map[string]string
	Headers       map[string]string
	Body          string
	TimeoutSecond int
}

type HttpClientRes struct {
	Status  string
	Code    int
	Headers map[string]string
	Body    string
}

func (a *HttpClientApp) DoHttp(req *HttpClientReq) (*HttpClientRes, error) {
	// 创建 HTTP 请求
	httpReq, err := http.NewRequest(req.Method, req.URL, strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	// 设置查询参数
	if req.Query != nil {
		query := httpReq.URL.Query()
		for key, value := range req.Query {
			query.Add(key, value)
		}
		httpReq.URL.RawQuery = query.Encode()
	}

	// 设置请求头
	if req.Headers != nil {
		for key, value := range req.Headers {
			httpReq.Header.Set(key, value)
		}
	}

	cli := http.Client{
		Timeout: time.Duration(req.TimeoutSecond) * time.Second,
	}

	// 发送 HTTP 请求
	httpRes, err := cli.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	// 创建 HTTP 响应
	res := &HttpClientRes{
		Status:  httpRes.Status,
		Code:    httpRes.StatusCode,
		Headers: make(map[string]string),
		Body:    string(body),
	}

	// 复制响应头
	for key, values := range httpRes.Header {
		res.Headers[key] = strings.Join(values, ",")
	}

	return res, nil
}
