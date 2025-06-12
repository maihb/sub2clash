package parser

import (
	"fmt"
	"net/url"
	"strings"

	P "github.com/bestnite/sub2clash/model/proxy"
)

type SocksParser struct{}

func (p *SocksParser) SupportClash() bool {
	return true
}
func (p *SocksParser) SupportMeta() bool {
	return true
}

func (p *SocksParser) GetPrefixes() []string {
	return []string{"socks://"}
}

func (p *SocksParser) GetType() string {
	return "socks5"
}

func (p *SocksParser) Parse(proxy string) (P.Proxy, error) {
	if !hasPrefix(proxy, p.GetPrefixes()) {
		return P.Proxy{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	link, err := url.Parse(proxy)
	if err != nil {
		return P.Proxy{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "url parse error",
			Raw:     proxy,
		}
	}
	server := link.Hostname()
	if server == "" {
		return P.Proxy{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host",
			Raw:     proxy,
		}
	}
	portStr := link.Port()
	if portStr == "" {
		return P.Proxy{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server port",
			Raw:     proxy,
		}
	}
	port, err := ParsePort(portStr)
	if err != nil {
		return P.Proxy{}, &ParseError{
			Type: ErrInvalidPort,
			Raw:  portStr,
		}
	}

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	encodeStr := link.User.Username()
	var username, password string
	if encodeStr != "" {
		decodeStr, err := DecodeBase64(encodeStr)
		splitStr := strings.Split(decodeStr, ":")
		if err != nil {
			return P.Proxy{}, &ParseError{
				Type:    ErrInvalidStruct,
				Message: "url parse error",
				Raw:     proxy,
			}
		}
		username = splitStr[0]
		if len(splitStr) == 2 {
			password = splitStr[1]
		}
	}
	return P.Proxy{
		Type: p.GetType(),
		Name: remarks,
		Socks: P.Socks{
			Server:   server,
			Port:     port,
			UserName: username,
			Password: password,
		},
	}, nil

}

func init() {
	RegisterParser(&SocksParser{})
}
