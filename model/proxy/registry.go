package proxy

import (
	"sync"
)

// ProxyMarshaler 代理YAML序列化接口
type ProxyMarshaler interface {
	// MarshalProxy 将通用Proxy对象序列化为特定协议的YAML结构
	MarshalProxy(p Proxy) (interface{}, error)
	// GetType 返回支持的协议类型
	GetType() string
}

// marshalerRegistry YAML序列化器注册中心
type marshalerRegistry struct {
	mu         sync.RWMutex
	marshalers map[string]ProxyMarshaler // type -> marshaler
}

var yamlRegistry = &marshalerRegistry{
	marshalers: make(map[string]ProxyMarshaler),
}

// RegisterMarshaler 注册YAML序列化器
func RegisterMarshaler(marshaler ProxyMarshaler) {
	yamlRegistry.mu.Lock()
	defer yamlRegistry.mu.Unlock()

	yamlRegistry.marshalers[marshaler.GetType()] = marshaler
}

// GetMarshaler 根据协议类型获取序列化器
func GetMarshaler(proxyType string) (ProxyMarshaler, bool) {
	yamlRegistry.mu.RLock()
	defer yamlRegistry.mu.RUnlock()

	marshaler, exists := yamlRegistry.marshalers[proxyType]
	return marshaler, exists
}

// GetAllMarshalers 获取所有注册的序列化器
func GetAllMarshalers() map[string]ProxyMarshaler {
	yamlRegistry.mu.RLock()
	defer yamlRegistry.mu.RUnlock()

	result := make(map[string]ProxyMarshaler)
	for k, v := range yamlRegistry.marshalers {
		result[k] = v
	}
	return result
}

// GetSupportedTypes 获取所有支持的协议类型
func GetSupportedTypes() []string {
	yamlRegistry.mu.RLock()
	defer yamlRegistry.mu.RUnlock()

	types := make([]string, 0, len(yamlRegistry.marshalers))
	for proxyType := range yamlRegistry.marshalers {
		types = append(types, proxyType)
	}
	return types
}
