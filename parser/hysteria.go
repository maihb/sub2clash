package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	E "github.com/bestnite/sub2clash/error"
	P "github.com/bestnite/sub2clash/model/proxy"
)

type HysteriaParser struct{}

func (p *HysteriaParser) SupportClash() bool {
	return false
}

func (p *HysteriaParser) SupportMeta() bool {
	return true
}

func (p *HysteriaParser) GetPrefixes() []string {
	return []string{"hysteria://"}
}

func (p *HysteriaParser) GetType() string {
	return "hysteria"
}

func (p *HysteriaParser) Parse(proxy string) (P.Proxy, error) {
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
			Type:    E.ErrInvalidPort,
			Message: err.Error(),
			Raw:     proxy,
		}
	}

	query := link.Query()

	protocol, auth, insecure, upmbps, downmbps, obfs, alpnStr := query.Get("protocol"), query.Get("auth"), query.Get("insecure"), query.Get("upmbps"), query.Get("downmbps"), query.Get("obfs"), query.Get("alpn")
	insecureBool, err := strconv.ParseBool(insecure)
	if err != nil {
		insecureBool = false
	}

	var alpn []string
	alpnStr = strings.TrimSpace(alpnStr)
	if alpnStr != "" {
		alpn = strings.Split(alpnStr, ",")
	}

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := P.Proxy{
		Type: p.GetType(),
		Name: remarks,
		Hysteria: P.Hysteria{
			Server:         server,
			Port:           port,
			Up:             upmbps,
			Down:           downmbps,
			Auth:           auth,
			Obfs:           obfs,
			SkipCertVerify: insecureBool,
			Alpn:           alpn,
			Protocol:       protocol,
			AllowInsecure:  insecureBool,
		},
	}
	return result, nil
}

func init() {
	RegisterParser(&HysteriaParser{})
}
