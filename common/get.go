package common

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nitezs/sub2clash/config"
)

type GetConfig struct {
	userAgent string
}

type GetOption func(*GetConfig)

func WithUserAgent(userAgent string) GetOption {
	return func(config *GetConfig) {
		config.userAgent = userAgent
	}
}

func Get(url string, options ...GetOption) (resp *http.Response, err error) {
	retryTimes := config.Default.RequestRetryTimes
	haveTried := 0
	retryDelay := time.Second
	getConfig := GetConfig{}
	for _, option := range options {
		option(&getConfig)
	}
	var req *http.Request
	var get *http.Response
	for haveTried < retryTimes {
		client := &http.Client{}
		//client.Timeout = time.Second * 10
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			haveTried++
			time.Sleep(retryDelay)
			continue
		}
		if getConfig.userAgent != "" {
			req.Header.Set("User-Agent", getConfig.userAgent)
		}
		get, err = client.Do(req)
		if err != nil {
			haveTried++
			time.Sleep(retryDelay)
			continue
		} else {
			if get != nil && get.ContentLength > config.Default.RequestMaxFileSize {
				return nil, errors.New("文件过大")
			}
			return get, nil
		}

	}
	return nil, fmt.Errorf("请求失败：%v", err)
}

func Head(url string, options ...GetOption) (resp *http.Response, err error) {
	retryTimes := config.Default.RequestRetryTimes
	haveTried := 0
	retryDelay := time.Second

	// 解析可选参数（如 User-Agent）
	getConfig := GetConfig{}
	for _, option := range options {
		option(&getConfig)
	}

	var req *http.Request
	var headResp *http.Response

	for haveTried < retryTimes {
		client := &http.Client{}
		req, err = http.NewRequest("HEAD", url, nil)
		if err != nil {
			haveTried++
			time.Sleep(retryDelay)
			continue
		}

		// 设置 User-Agent（如果提供）
		if getConfig.userAgent != "" {
			req.Header.Set("User-Agent", getConfig.userAgent)
		}

		headResp, err = client.Do(req)
		if err != nil {
			haveTried++
			time.Sleep(retryDelay)
			continue
		}

		// HEAD 请求不检查 ContentLength，因为没有响应体
		return headResp, nil
	}

	return nil, fmt.Errorf("HEAD 请求失败：%v", err)
}
