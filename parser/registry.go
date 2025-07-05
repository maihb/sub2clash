package parser

import (
	"fmt"
	"strings"
	"sync"

	P "github.com/bestnite/sub2clash/model/proxy"
)

type ProxyParser interface {
	Parse(config ParseConfig, proxy string) (P.Proxy, error)
	GetPrefixes() []string
	GetType() string
	SupportClash() bool
	SupportMeta() bool
}

type parserRegistry struct {
	mu      sync.RWMutex
	parsers map[string]ProxyParser
}

var registry = &parserRegistry{
	parsers: make(map[string]ProxyParser),
}

func RegisterParser(parser ProxyParser) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	for _, prefix := range parser.GetPrefixes() {
		registry.parsers[prefix] = parser
	}
}

func GetParser(prefix string) (ProxyParser, bool) {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	parser, exists := registry.parsers[prefix]
	return parser, exists
}

func GetAllParsers() map[string]ProxyParser {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	result := make(map[string]ProxyParser)
	for k, v := range registry.parsers {
		result[k] = v
	}
	return result
}

func GetAllPrefixes() []string {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	prefixes := make([]string, 0, len(registry.parsers))
	for prefix := range registry.parsers {
		prefixes = append(prefixes, prefix)
	}
	return prefixes
}

func ParseProxyWithRegistry(config ParseConfig, proxy string) (P.Proxy, error) {
	proxy = strings.TrimSpace(proxy)
	if proxy == "" {
		return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidStruct, "empty proxy string")
	}

	for prefix, parser := range registry.parsers {
		if strings.HasPrefix(proxy, prefix) {
			return parser.Parse(config, proxy)
		}
	}

	return P.Proxy{}, fmt.Errorf("%w: %s", ErrInvalidPrefix, "unsupported protocol")
}
