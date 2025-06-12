package test

import (
	"errors"
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestShadowsocks_Basic_SimpleLink(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://YWVzLTI1Ni1nY206cGFzc3dvcmQ=@127.0.0.1:8080"

	expected := proxy.Proxy{
		Type: "ss",
		Name: "127.0.0.1:8080",
		ShadowSocks: proxy.ShadowSocks{
			Server:   "127.0.0.1",
			Port:     8080,
			Cipher:   "aes-256-gcm",
			Password: "password",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocks_Basic_IPv6Address(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://YWVzLTI1Ni1nY206cGFzc3dvcmQ=@[2001:db8::1]:8080"

	expected := proxy.Proxy{
		Type: "ss",
		Name: "2001:db8::1:8080",
		ShadowSocks: proxy.ShadowSocks{
			Server:   "2001:db8::1",
			Port:     8080,
			Cipher:   "aes-256-gcm",
			Password: "password",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocks_Basic_WithRemark(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://YWVzLTI1Ni1nY206cGFzc3dvcmQ=@proxy.example.com:8080#My%20SS%20Proxy"

	expected := proxy.Proxy{
		Type: "ss",
		Name: "My SS Proxy",
		ShadowSocks: proxy.ShadowSocks{
			Server:   "proxy.example.com",
			Port:     8080,
			Cipher:   "aes-256-gcm",
			Password: "password",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocks_Advanced_Base64FullEncoded(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://YWVzLTI1Ni1nY206cGFzc3dvcmRAbG9jYWxob3N0OjgwODA=#Local%20SS"

	expected := proxy.Proxy{
		Type: "ss",
		Name: "Local SS",
		ShadowSocks: proxy.ShadowSocks{
			Server:   "localhost",
			Port:     8080,
			Cipher:   "aes-256-gcm",
			Password: "password",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocks_Advanced_PlainUserPassword(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://aes-256-gcm:password@192.168.1.1:8080"

	expected := proxy.Proxy{
		Type: "ss",
		Name: "192.168.1.1:8080",
		ShadowSocks: proxy.ShadowSocks{
			Server:   "192.168.1.1",
			Port:     8080,
			Cipher:   "aes-256-gcm",
			Password: "password",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocks_Advanced_ChaCha20Cipher(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://chacha20-poly1305:mypassword@server.com:443#ChaCha20"

	expected := proxy.Proxy{
		Type: "ss",
		Name: "ChaCha20",
		ShadowSocks: proxy.ShadowSocks{
			Server:   "server.com",
			Port:     443,
			Cipher:   "chacha20-poly1305",
			Password: "mypassword",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

// 错误处理测试
func TestShadowsocks_Error_MissingServer(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://YWVzLTI1Ni1nY206cGFzc3dvcmQ=@:8080"

	_, err := p.Parse(input)
	if !errors.Is(err, parser.ErrInvalidStruct) {
		t.Errorf("Error is not expected: %v", err)
	}
}

func TestShadowsocks_Error_MissingPort(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "ss://YWVzLTI1Ni1nY206cGFzc3dvcmQ=@127.0.0.1"

	_, err := p.Parse(input)
	if !errors.Is(err, parser.ErrInvalidStruct) {
		t.Errorf("Error is not expected: %v", err)
	}
}

func TestShadowsocks_Error_InvalidProtocol(t *testing.T) {
	p := &parser.ShadowsocksParser{}
	input := "http://example.com:8080"

	_, err := p.Parse(input)
	if !errors.Is(err, parser.ErrInvalidPrefix) {
		t.Errorf("Error is not expected: %v", err)
	}
}
