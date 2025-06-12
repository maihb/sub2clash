package parser

import (
	"strings"
	"sync"

	P "github.com/bestnite/sub2clash/model/proxy"
)

// ProxyParser 定义代理解析器接口
type ProxyParser interface {
	// Parse 解析代理字符串并返回Proxy对象
	Parse(proxy string) (P.Proxy, error)
	// GetPrefix 返回支持的协议前缀（可以返回多个）
	GetPrefixes() []string
	// GetType 返回协议类型名称
	GetType() string
	SupportClash() bool
	SupportMeta() bool
}

// parserRegistry 解析器注册中心
type parserRegistry struct {
	mu      sync.RWMutex
	parsers map[string]ProxyParser // prefix -> parser
}

var registry = &parserRegistry{
	parsers: make(map[string]ProxyParser),
}

// RegisterParser 注册解析器
func RegisterParser(parser ProxyParser) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	for _, prefix := range parser.GetPrefixes() {
		registry.parsers[prefix] = parser
	}
}

// GetParser 根据前缀获取解析器
func GetParser(prefix string) (ProxyParser, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	parser, exists := registry.parsers[prefix]
	return parser, exists
}

// GetAllParsers 获取所有注册的解析器
func GetAllParsers() map[string]ProxyParser {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	result := make(map[string]ProxyParser)
	for k, v := range registry.parsers {
		result[k] = v
	}
	return result
}

// ParseProxyWithRegistry 使用注册机制解析代理
func ParseProxyWithRegistry(proxy string) (P.Proxy, error) {
	proxy = strings.TrimSpace(proxy)
	if proxy == "" {
		return P.Proxy{}, &ParseError{Type: ErrInvalidStruct, Raw: proxy, Message: "empty proxy string"}
	}

	// 查找匹配的解析器
	for prefix, parser := range registry.parsers {
		if strings.HasPrefix(proxy, prefix) {
			return parser.Parse(proxy)
		}
	}

	return P.Proxy{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy, Message: "unsupported protocol"}
}
