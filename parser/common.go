package parser

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	P "github.com/bestnite/sub2clash/model/proxy"
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

func isLikelyBase64(s string) bool {
	if len(s)%4 == 0 && strings.HasSuffix(s, "=") && !strings.Contains(strings.TrimSuffix(s, "="), "=") {
		s = strings.TrimSuffix(s, "=")
		chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
		for _, c := range s {
			if !strings.ContainsRune(chars, c) {
				return false
			}
		}
		return true
	}
	return false
}

func DecodeBase64(s string) (string, error) {
	s = strings.TrimSpace(s)

	if strings.Contains(s, "-") || strings.Contains(s, "_") {
		s = strings.Replace(s, "-", "+", -1)
		s = strings.Replace(s, "_", "/", -1)
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

func ParseProxies(proxies ...string) ([]P.Proxy, error) {
	var result []P.Proxy
	for _, proxy := range proxies {
		if proxy != "" {
			var proxyItem P.Proxy
			var err error

			proxyItem, err = ParseProxyWithRegistry(proxy)
			if err != nil {
				return nil, err
			}
			result = append(result, proxyItem)
		}
	}
	return result, nil
}
