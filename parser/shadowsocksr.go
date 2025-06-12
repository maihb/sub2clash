package parser

import (
	"net/url"
	"strconv"
	"strings"

	P "github.com/bestnite/sub2clash/model/proxy"
)

type ShadowsocksRParser struct{}

func (p *ShadowsocksRParser) SupportClash() bool {
	return true
}

func (p *ShadowsocksRParser) SupportMeta() bool {
	return true
}

func (p *ShadowsocksRParser) GetPrefixes() []string {
	return []string{"ssr://"}
}

func (p *ShadowsocksRParser) GetType() string {
	return "ssr"
}

func (p *ShadowsocksRParser) Parse(proxy string) (P.Proxy, error) {
	if !hasPrefix(proxy, p.GetPrefixes()) {
		return P.Proxy{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	for _, prefix := range p.GetPrefixes() {
		if strings.HasPrefix(proxy, prefix) {
			proxy = strings.TrimPrefix(proxy, prefix)
			break
		}
	}

	proxy, err := DecodeBase64(proxy)
	if err != nil {
		return P.Proxy{}, &ParseError{
			Type: ErrInvalidBase64,
			Raw:  proxy,
		}
	}
	serverInfoAndParams := strings.SplitN(proxy, "/?", 2)
	parts := strings.Split(serverInfoAndParams[0], ":")
	server := parts[0]
	protocol := parts[2]
	method := parts[3]
	obfs := parts[4]
	password, err := DecodeBase64(parts[5])
	if err != nil {
		return P.Proxy{}, &ParseError{
			Type:    ErrInvalidStruct,
			Raw:     proxy,
			Message: err.Error(),
		}
	}
	port, err := ParsePort(parts[1])
	if err != nil {
		return P.Proxy{}, &ParseError{
			Type:    ErrInvalidPort,
			Message: err.Error(),
			Raw:     proxy,
		}
	}

	var obfsParam string
	var protoParam string
	var remarks string
	if len(serverInfoAndParams) == 2 {
		params, err := url.ParseQuery(serverInfoAndParams[1])
		if err != nil {
			return P.Proxy{}, &ParseError{
				Type:    ErrCannotParseParams,
				Raw:     proxy,
				Message: err.Error(),
			}
		}
		if params.Get("obfsparam") != "" {
			obfsParam, err = DecodeBase64(params.Get("obfsparam"))
		}
		if params.Get("protoparam") != "" {
			protoParam, err = DecodeBase64(params.Get("protoparam"))
		}
		if params.Get("remarks") != "" {
			remarks, err = DecodeBase64(params.Get("remarks"))
		} else {
			remarks = server + ":" + strconv.Itoa(port)
		}
		if err != nil {
			return P.Proxy{}, &ParseError{
				Type:    ErrInvalidStruct,
				Raw:     proxy,
				Message: err.Error(),
			}
		}
	}

	result := P.Proxy{
		Type: p.GetType(),
		Name: remarks,
		ShadowSocksR: P.ShadowSocksR{
			Server:        server,
			Port:          port,
			Protocol:      protocol,
			Cipher:        method,
			Obfs:          obfs,
			Password:      password,
			ObfsParam:     obfsParam,
			ProtocolParam: protoParam,
		},
	}

	return result, nil
}

func init() {
	RegisterParser(&ShadowsocksRParser{})
}
