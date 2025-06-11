package parser

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	E "github.com/bestnite/sub2clash/error"
	P "github.com/bestnite/sub2clash/model/proxy"
)

type VmessParser struct{}

func (p *VmessParser) GetPrefixes() []string {
	return []string{"vmess://"}
}

func (p *VmessParser) GetType() string {
	return "vmess"
}

func (p *VmessParser) Parse(proxy string) (P.Proxy, error) {
	if !hasPrefix(proxy, p.GetPrefixes()) {
		return P.Proxy{}, &E.ParseError{Type: E.ErrInvalidPrefix, Raw: proxy}
	}

	for _, prefix := range p.GetPrefixes() {
		if strings.HasPrefix(proxy, prefix) {
			proxy = strings.TrimPrefix(proxy, prefix)
			break
		}
	}
	base64, err := DecodeBase64(proxy)
	if err != nil {
		return P.Proxy{}, &E.ParseError{Type: E.ErrInvalidBase64, Raw: proxy, Message: err.Error()}
	}

	var vmess P.VmessJson
	err = json.Unmarshal([]byte(base64), &vmess)
	if err != nil {
		return P.Proxy{}, &E.ParseError{Type: E.ErrInvalidStruct, Raw: proxy, Message: err.Error()}
	}

	var port int
	switch vmess.Port.(type) {
	case string:
		port, err = ParsePort(vmess.Port.(string))
		if err != nil {
			return P.Proxy{}, &E.ParseError{
				Type:    E.ErrInvalidPort,
				Message: err.Error(),
				Raw:     proxy,
			}
		}
	case float64:
		port = int(vmess.Port.(float64))
	}

	aid := 0
	switch vmess.Aid.(type) {
	case string:
		aid, err = strconv.Atoi(vmess.Aid.(string))
		if err != nil {
			return P.Proxy{}, &E.ParseError{Type: E.ErrInvalidStruct, Raw: proxy, Message: err.Error()}
		}
	case float64:
		aid = int(vmess.Aid.(float64))
	}

	if vmess.Scy == "" {
		vmess.Scy = "auto"
	}

	name, err := url.QueryUnescape(vmess.Ps)
	if err != nil {
		name = vmess.Ps
	}

	result := P.Proxy{
		Name:    name,
		Type:    p.GetType(),
		Server:  vmess.Add,
		Port:    port,
		UUID:    vmess.Id,
		AlterID: aid,
		Cipher:  vmess.Scy,
	}

	if vmess.Tls == "tls" {
		var alpn []string
		if strings.Contains(vmess.Alpn, ",") {
			alpn = strings.Split(vmess.Alpn, ",")
		} else {
			alpn = nil
		}
		result.TLS = true
		result.Fingerprint = vmess.Fp
		result.Alpn = alpn
		result.Servername = vmess.Sni
	}

	if vmess.Net == "ws" {
		if vmess.Path == "" {
			vmess.Path = "/"
		}
		if vmess.Host == "" {
			vmess.Host = vmess.Add
		}
		result.Network = "ws"
		result.WSOpts = P.WSOptions{
			Path: vmess.Path,
			Headers: map[string]string{
				"Host": vmess.Host,
			},
		}
	}

	if vmess.Net == "grpc" {
		result.GrpcOpts = P.GrpcOptions{
			GrpcServiceName: vmess.Path,
		}
		result.Network = "grpc"
	}

	if vmess.Net == "h2" {
		result.HTTP2Opts = P.HTTP2Options{
			Host: strings.Split(vmess.Host, ","),
			Path: vmess.Path,
		}
		result.Network = "h2"
	}

	return result, nil
}

func init() {
	RegisterParser(&VmessParser{})
}
