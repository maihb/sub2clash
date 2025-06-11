package parser

import (
	"fmt"
	"net/url"
	"strings"

	E "github.com/bestnite/sub2clash/error"
	P "github.com/bestnite/sub2clash/model/proxy"
)

type Hysteria2Parser struct{}

func (p *Hysteria2Parser) SupportClash() bool {
	return false
}

func (p *Hysteria2Parser) SupportMeta() bool {
	return true
}

func (p *Hysteria2Parser) GetPrefixes() []string {
	return []string{"hysteria2://", "hy2://"}
}

func (p *Hysteria2Parser) GetType() string {
	return "hysteria2"
}

func (p *Hysteria2Parser) Parse(proxy string) (P.Proxy, error) {
	if !hasPrefix(proxy, p.GetPrefixes()) {
		return P.Proxy{}, &E.ParseError{Type: E.ErrInvalidPrefix, Raw: proxy}
	}

	link, err := url.Parse(proxy)
	if err != nil {
		return P.Proxy{}, &E.ParseError{
			Type:    E.ErrInvalidStruct,
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
			Type: E.ErrInvalidPort,
			Raw:  portStr,
		}
	}
	network, obfs, obfsPassword, pinSHA256, insecure, sni := query.Get("network"), query.Get("obfs"), query.Get("obfs-password"), query.Get("pinSHA256"), query.Get("insecure"), query.Get("sni")
	enableTLS := pinSHA256 != "" || sni != ""
	insecureBool := insecure == "1"
	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := P.Proxy{
		Type:           p.GetType(),
		Name:           remarks,
		Server:         server,
		Port:           port,
		Password:       password,
		Obfs:           obfs,
		ObfsParam:      obfsPassword,
		Sni:            sni,
		SkipCertVerify: insecureBool,
		TLS:            enableTLS,
		Network:        network,
	}
	return result, nil
}

func init() {
	RegisterParser(&Hysteria2Parser{})
}
