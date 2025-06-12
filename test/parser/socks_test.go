package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestSocks_Basic_SimpleLink(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@127.0.0.1:1080#SOCKS%20Proxy"

	expected := proxy.Proxy{
		Type: "socks5",
		Name: "SOCKS Proxy",
		Socks: proxy.Socks{
			Server:   "127.0.0.1",
			Port:     1080,
			UserName: "user",
			Password: "pass",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestSocks_Basic_NoAuth(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://127.0.0.1:1080#SOCKS%20No%20Auth"

	expected := proxy.Proxy{
		Type: "socks5",
		Name: "SOCKS No Auth",
		Socks: proxy.Socks{
			Server: "127.0.0.1",
			Port:   1080,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestSocks_Basic_IPv6Address(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@[2001:db8::1]:1080#SOCKS%20IPv6"

	expected := proxy.Proxy{
		Type: "socks5",
		Name: "SOCKS IPv6",
		Socks: proxy.Socks{
			Server:   "2001:db8::1",
			Port:     1080,
			UserName: "user",
			Password: "pass",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestSocks_Basic_WithTLS(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@127.0.0.1:1080?tls=true&sni=example.com#SOCKS%20TLS"

	expected := proxy.Proxy{
		Type: "socks5",
		Name: "SOCKS TLS",
		Socks: proxy.Socks{
			Server:   "127.0.0.1",
			Port:     1080,
			UserName: "user",
			Password: "pass",
			TLS:      true,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestSocks_Basic_WithUDP(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@127.0.0.1:1080?udp=true#SOCKS%20UDP"

	expected := proxy.Proxy{
		Type: "socks5",
		Name: "SOCKS UDP",
		Socks: proxy.Socks{
			Server:   "127.0.0.1",
			Port:     1080,
			UserName: "user",
			Password: "pass",
			UDP:      true,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestSocks_Error_MissingServer(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@:1080"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestSocks_Error_MissingPort(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@127.0.0.1"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestSocks_Error_InvalidPort(t *testing.T) {
	p := &parser.SocksParser{}
	input := "socks://user:pass@127.0.0.1:99999"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestSocks_Error_InvalidProtocol(t *testing.T) {
	p := &parser.SocksParser{}
	input := "ss://example.com:8080"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
