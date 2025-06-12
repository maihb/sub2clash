package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestHysteria_Basic_SimpleLink(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://127.0.0.1:8080?protocol=udp&auth=password123&upmbps=100&downmbps=100#Hysteria%20Proxy"

	expected := proxy.Proxy{
		Type: "hysteria",
		Name: "Hysteria Proxy",
		Hysteria: proxy.Hysteria{
			Server:         "127.0.0.1",
			Port:           8080,
			Protocol:       "udp",
			Auth:           "password123",
			Up:             "100",
			Down:           "100",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria_Basic_WithAuthString(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://proxy.example.com:443?protocol=wechat-video&auth-str=myauth&upmbps=50&downmbps=200&insecure=true#Hysteria%20Auth"

	expected := proxy.Proxy{
		Type: "hysteria",
		Name: "Hysteria Auth",
		Hysteria: proxy.Hysteria{
			Server:         "proxy.example.com",
			Port:           443,
			Protocol:       "wechat-video",
			AuthString:     "myauth",
			Up:             "50",
			Down:           "200",
			SkipCertVerify: true,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria_Basic_WithObfs(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://127.0.0.1:8080?auth=password123&upmbps=100&downmbps=100&obfs=xplus&alpn=h3#Hysteria%20Obfs"

	expected := proxy.Proxy{
		Type: "hysteria",
		Name: "Hysteria Obfs",
		Hysteria: proxy.Hysteria{
			Server:         "127.0.0.1",
			Port:           8080,
			Auth:           "password123",
			Up:             "100",
			Down:           "100",
			Obfs:           "xplus",
			ALPN:           []string{"h3"},
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria_Basic_IPv6Address(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://[2001:db8::1]:8080?auth=password123&upmbps=100&downmbps=100#Hysteria%20IPv6"

	expected := proxy.Proxy{
		Type: "hysteria",
		Name: "Hysteria IPv6",
		Hysteria: proxy.Hysteria{
			Server:         "2001:db8::1",
			Port:           8080,
			Auth:           "password123",
			Up:             "100",
			Down:           "100",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria_Basic_MultiALPN(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://proxy.example.com:443?auth=password123&upmbps=100&downmbps=100&alpn=h3,h2,http/1.1#Hysteria%20Multi%20ALPN"

	expected := proxy.Proxy{
		Type: "hysteria",
		Name: "Hysteria Multi ALPN",
		Hysteria: proxy.Hysteria{
			Server:         "proxy.example.com",
			Port:           443,
			Auth:           "password123",
			Up:             "100",
			Down:           "100",
			ALPN:           []string{"h3", "h2", "http/1.1"},
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestHysteria_Error_MissingServer(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://:8080?auth=password123"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria_Error_MissingPort(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://127.0.0.1?auth=password123"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria_Error_InvalidPort(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria://127.0.0.1:99999?auth=password123"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestHysteria_Error_InvalidProtocol(t *testing.T) {
	p := &parser.HysteriaParser{}
	input := "hysteria2://example.com:8080"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
