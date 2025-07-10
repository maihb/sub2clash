package parser

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"

	P "github.com/maihb/sub2clash/model/proxy"
)

func hasPrefix(proxy string, prefixes []string) bool {
	hasPrefix := false
	for _, prefix := range prefixes {
		if strings.HasPrefix(proxy, prefix) {
			hasPrefix = true
			break
		}
	}
	return hasPrefix
}

func ParsePort(portStr string) (int, error) {
	port, err := strconv.Atoi(portStr)

	if err != nil {
		return 0, err
	}
	if port < 1 || port > 65535 {
		return 0, errors.New("invaild port range")
	}
	return port, nil
}

// isLikelyBase64 不严格判断是否是合法的 Base64, 很多分享链接不符合 Base64 规范
func isLikelyBase64(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}

	if !strings.Contains(strings.TrimSuffix(s, "="), "=") {
		s = strings.TrimSuffix(s, "=")
		chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
		for _, c := range s {
			if !strings.ContainsRune(chars, c) {
				return false
			}
		}
	}

	decoded, err := DecodeBase64(s)
	if err != nil {
		return false
	}
	if !utf8.ValidString(decoded) {
		return false
	}

	return true
}

func DecodeBase64(s string) (string, error) {
	s = strings.TrimSpace(s)

	if strings.Contains(s, "-") || strings.Contains(s, "_") {
		s = strings.ReplaceAll(s, "-", "+")
		s = strings.ReplaceAll(s, "_", "/")
	}
	if len(s)%4 != 0 {
		s += strings.Repeat("=", 4-len(s)%4)
	}
	decodeStr, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decodeStr), nil
}

type ParseConfig struct {
	UseUDP bool
}

func ParseProxies(config ParseConfig, proxies ...string) ([]P.Proxy, error) {
	var result []P.Proxy
	for _, proxy := range proxies {
		if proxy != "" {
			var proxyItem P.Proxy
			var err error

			proxyItem, err = ParseProxyWithRegistry(config, proxy)
			if err != nil {
				return nil, err
			}
			result = append(result, proxyItem)
		}
	}
	return result, nil
}
