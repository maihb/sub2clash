package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nitezs/sub2clash/constant"
	"github.com/nitezs/sub2clash/model"
)

func ParseAnytls(proxy string) (model.Proxy, error) {
	if !strings.HasPrefix(proxy, constant.AnytlsPrefix) {
		return model.Proxy{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	link, err := url.Parse(proxy)
	if err != nil {
		return model.Proxy{}, &ParseError{
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
		return model.Proxy{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host",
			Raw:     proxy,
		}
	}
	portStr := link.Port()
	if portStr == "" {
		return model.Proxy{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server port",
			Raw:     proxy,
		}
	}
	port, err := ParsePort(portStr)
	if err != nil {
		return model.Proxy{}, &ParseError{
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

	result := model.Proxy{
		Type:           "anytls",
		Name:           remarks,
		Server:         server,
		Port:           port,
		Password:       password,
		Sni:            sni,
		SkipCertVerify: insecureBool,
	}
	return result, nil
}
