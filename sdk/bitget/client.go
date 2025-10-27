package bitget

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/time/rate"
)

const (
	BaseURL = "https://bopenapi.bgwapi.io"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type BClient struct {
	apiKey    string
	apiSecret string
	BaseURL   string
	limiter   *rate.Limiter
	// rdb       *redis.Client
	http *http.Client
}

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewClient(apiKey, apiSecret string, timeOut time.Duration) *BClient {
	return &BClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		BaseURL:   BaseURL,
		limiter:   rate.NewLimiter(rate.Limit(100), 1),
		http: &http.Client{
			Timeout: timeOut,
		},
	}
}

func (c *BClient) createSignature(path string, query map[string]string, body string) (string, string) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	contentMap := make(map[string]string)
	contentMap["x-api-key"] = c.apiKey
	contentMap["x-api-timestamp"] = timestamp
	contentMap["apiPath"] = path
	for key, value := range query {
		contentMap[key] = value
	}
	contentMap["body"] = body
	content, _ := json.Marshal(contentMap)
	mac := hmac.New(sha256.New, []byte(c.apiSecret))
	mac.Write(content)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), timestamp
}

func (c *BClient) sendGetRequest(ctx context.Context, requestPath string, params map[string]interface{}) ([]byte, error) {
	// 生成签名
	strParams := make(map[string]string)
	for k, v := range params {
		strParams[k] = fmt.Sprintf("%v", v)
	}
	signature, timestamp := c.createSignature(requestPath, strParams, "")

	// 生成请求头
	headers := map[string]string{
		"x-api-key":       c.apiKey,
		"x-api-signature": signature,
		"x-api-timestamp": timestamp,
	}
	// 构建查询参数
	query := url.Values{}

	for k, v := range strParams {
		query.Add(k, v)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "GET", BaseURL+requestPath+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("request", req.URL.String())

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err := c.limiter.Wait(ctx); err != nil {
		return nil, err
	}

	// 发送请求
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sendGetRequest failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sendGetRequest failed: %v", resp.Status)
	}

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *BClient) sendPostRequest(ctx context.Context, requestPath string, params interface{}) ([]byte, error) {
	// 准备请求体
	var reqBody []byte
	if params != nil {
		reqBody, _ = json.Marshal(params)
	}
	// 生成签名
	signature, timestamp := c.createSignature(requestPath, nil, string(reqBody))

	// 生成请求头
	headers := map[string]string{
		"x-api-key":       c.apiKey,
		"x-api-signature": signature,
		"x-api-timestamp": timestamp,
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "POST", BaseURL+requestPath, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	fmt.Println("request", req.URL.String())

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sendPostRequest failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sendPostRequest failed: %v", resp.Status)
	}

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
