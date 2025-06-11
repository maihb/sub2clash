package parser

import (
	"fmt"
	"net/url"
	"strings"

	E "github.com/bestnite/sub2clash/error"
	P "github.com/bestnite/sub2clash/model/proxy"
)

// ShadowsocksParser Shadowsocks协议解析器
type ShadowsocksParser struct{}

func (p *ShadowsocksParser) SupportClash() bool {
	return true
}

func (p *ShadowsocksParser) SupportMeta() bool {
	return true
}

// GetPrefixes 返回支持的协议前缀
func (p *ShadowsocksParser) GetPrefixes() []string {
	return []string{"ss://"}
}

// GetType 返回协议类型
func (p *ShadowsocksParser) GetType() string {
	return "ss"
}

// Parse 解析Shadowsocks代理
func (p *ShadowsocksParser) Parse(proxy string) (P.Proxy, error) {
	if !hasPrefix(proxy, p.GetPrefixes()) {
		return P.Proxy{}, &E.ParseError{Type: E.ErrInvalidPrefix, Raw: proxy}
	}

	if !strings.Contains(proxy, "@") {
		s := strings.SplitN(proxy, "#", 2)
		for _, prefix := range p.GetPrefixes() {
			if strings.HasPrefix(s[0], prefix) {
				s[0] = strings.TrimPrefix(s[0], prefix)
				break
			}
		}
		d, err := DecodeBase64(s[0])
		if err != nil {
			return P.Proxy{}, &E.ParseError{
				Type:    E.ErrInvalidStruct,
				Message: "url parse error",
				Raw:     proxy,
			}
		}
		if len(s) == 2 {
			proxy = "ss://" + d + "#" + s[1]
		} else {
			proxy = "ss://" + d
		}
	}
	link, err := url.Parse(proxy)
	if err != nil {
		return P.Proxy{}, &E.ParseError{
			Type:    E.ErrInvalidStruct,
			Message: "url parse error",
			Raw:     proxy,
		}
	}

	server := link.Hostname()
	if server == "" {
		return P.Proxy{}, &E.ParseError{
			Type:    E.ErrInvalidStruct,
			Message: "missing server host",
			Raw:     proxy,
		}
	}

	portStr := link.Port()
	if portStr == "" {
		return P.Proxy{}, &E.ParseError{
			Type:    E.ErrInvalidStruct,
			Message: "missing server port",
			Raw:     proxy,
		}
	}
	port, err := ParsePort(portStr)
	if err != nil {
		return P.Proxy{}, &E.ParseError{
			Type: E.ErrInvalidStruct,
			Raw:  proxy,
		}
	}

	method := link.User.Username()
	password, _ := link.User.Password()

	if password == "" {
		user, err := DecodeBase64(method)
		if err == nil {
			methodAndPass := strings.SplitN(user, ":", 2)
			if len(methodAndPass) == 2 {
				method = methodAndPass[0]
				password = methodAndPass[1]
			}
		}
	}
	if isLikelyBase64(password) {
		password, err = DecodeBase64(password)
		if err != nil {
			return P.Proxy{}, &E.ParseError{
				Type:    E.ErrInvalidStruct,
				Message: "password decode error",
				Raw:     proxy,
			}
		}
	}

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := P.Proxy{
		Type:     p.GetType(),
		Cipher:   method,
		Password: password,
		Server:   server,
		Port:     port,
		Name:     remarks,
	}

	return result, nil
}

// 注册解析器
func init() {
	RegisterParser(&ShadowsocksParser{})
}
