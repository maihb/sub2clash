package parser

import (
	"fmt"
	"net/url"
	"strings"

	P "github.com/bestnite/sub2clash/model/proxy"
)

type AnytlsParser struct{}

func (p *AnytlsParser) SupportClash() bool {
	return false
}

func (p *AnytlsParser) SupportMeta() bool {
	return true
}

func (p *AnytlsParser) GetPrefixes() []string {
	return []string{"anytls://"}
}

func (p *AnytlsParser) GetType() string {
	return "anytls"
}

func (p *AnytlsParser) Parse(proxy string) (P.Proxy, error) {
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

	username := link.User.Username()
	password, exist := link.User.Password()
	if !exist {
		password = username
	}

	query := link.Query()
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
	insecure, sni := query.Get("insecure"), query.Get("sni")
	insecureBool := insecure == "1"
	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := P.Proxy{
		Type: p.GetType(),
		Name: remarks,
		Anytls: P.Anytls{
			Server:         server,
			Port:           port,
			Password:       password,
			Sni:            sni,
			SkipCertVerify: insecureBool,
		},
	}
	return result, nil
}

func init() {
	RegisterParser(&AnytlsParser{})
}
