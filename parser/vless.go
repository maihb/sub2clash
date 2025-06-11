package parser

import (
	"fmt"
	"net/url"
	"strings"

	E "github.com/bestnite/sub2clash/error"
	P "github.com/bestnite/sub2clash/model/proxy"
)

type VlessParser struct{}

func (p *VlessParser) SupportClash() bool {
	return false
}

func (p *VlessParser) SupportMeta() bool {
	return true
}

func (p *VlessParser) GetPrefixes() []string {
	return []string{"vless://"}
}

func (p *VlessParser) GetType() string {
	return "vless"
}

func (p *VlessParser) Parse(proxy string) (P.Proxy, error) {
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
	port, err := ParsePort(portStr)
	if err != nil {
		return P.Proxy{}, &E.ParseError{
			Type:    E.ErrInvalidPort,
			Message: err.Error(),
			Raw:     proxy,
		}
	}

	query := link.Query()
	uuid := link.User.Username()
	flow, security, alpnStr, sni, insecure, fp, pbk, sid, path, host, serviceName, _type := query.Get("flow"), query.Get("security"), query.Get("alpn"), query.Get("sni"), query.Get("allowInsecure"), query.Get("fp"), query.Get("pbk"), query.Get("sid"), query.Get("path"), query.Get("host"), query.Get("serviceName"), query.Get("type")

	insecureBool := insecure == "1"
	var alpn []string
	if strings.Contains(alpnStr, ",") {
		alpn = strings.Split(alpnStr, ",")
	} else {
		alpn = nil
	}
	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := P.Proxy{
		Type:   p.GetType(),
		Server: server,
		Name:   remarks,
		Port:   port,
		UUID:   uuid,
		Flow:   flow,
	}

	if security == "tls" {
		result.TLS = true
		result.Alpn = alpn
		result.Sni = sni
		result.AllowInsecure = insecureBool
		result.ClientFingerprint = fp
	}

	if security == "reality" {
		result.TLS = true
		result.Servername = sni
		result.RealityOpts = P.RealityOptions{
			PublicKey: pbk,
			ShortID:   sid,
		}
		result.ClientFingerprint = fp
	}

	if _type == "ws" {
		result.Network = "ws"
		result.WSOpts = P.WSOptions{
			Path: path,
		}
		if host != "" {
			result.WSOpts.Headers = make(map[string]string)
			result.WSOpts.Headers["Host"] = host
		}
	}

	if _type == "grpc" {
		result.Network = "grpc"
		result.GrpcOpts = P.GrpcOptions{
			GrpcServiceName: serviceName,
		}
	}

	if _type == "http" {
		result.HTTPOpts = P.HTTPOptions{}
		result.HTTPOpts.Headers = map[string][]string{}

		result.HTTPOpts.Path = strings.Split(path, ",")

		hosts, err := url.QueryUnescape(host)
		if err != nil {
			return P.Proxy{}, &E.ParseError{
				Type:    E.ErrCannotParseParams,
				Raw:     proxy,
				Message: err.Error(),
			}
		}
		result.Network = "http"
		if hosts != "" {
			result.HTTPOpts.Headers["host"] = strings.Split(host, ",")
		}
	}

	return result, nil
}

func init() {
	RegisterParser(&VlessParser{})
}
