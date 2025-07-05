package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestShadowsocksR_Basic_SimpleLink(t *testing.T) {
	p := &parser.ShadowsocksRParser{}
	input := "ssr://MTI3LjAuMC4xOjQ0MzpvcmlnaW46YWVzLTE5Mi1jZmI6cGxhaW46TVRJek1USXovP2dyb3VwPVpHVm1ZWFZzZEEmcmVtYXJrcz1TRUZJUVE"

	expected := proxy.Proxy{
		Type: "ssr",
		Name: "HAHA",
		ShadowSocksR: proxy.ShadowSocksR{
			Server:        "127.0.0.1",
			Port:          443,
			Cipher:        "aes-192-cfb",
			Password:      "123123",
			ObfsParam:     "",
			Obfs:          "plain",
			Protocol:      "origin",
			ProtocolParam: "",
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocksR_Basic_WithParams(t *testing.T) {
	p := &parser.ShadowsocksRParser{}
	input := "ssr://MTI3LjAuMC4xOjQ0MzpvcmlnaW46YWVzLTE5Mi1jZmI6dGxzMS4wX3Nlc3Npb25fYXV0aDpNVEl6TVRJei8/b2Jmc3BhcmFtPWIySm1jeTF3WVhKaGJXVjBaWEkmcHJvdG9wYXJhbT1jSEp2ZEc5allXd3RjR0Z5WVcxbGRHVnkmZ3JvdXA9WkdWbVlYVnNkQSZyZW1hcmtzPVNFRklRUQ"

	expected := proxy.Proxy{
		Type: "ssr",
		Name: "HAHA",
		ShadowSocksR: proxy.ShadowSocksR{
			Server:        "127.0.0.1",
			Port:          443,
			Cipher:        "aes-192-cfb",
			Password:      "123123",
			ObfsParam:     "obfs-parameter",
			Obfs:          "tls1.0_session_auth",
			Protocol:      "origin",
			ProtocolParam: "protocal-parameter",
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocksR_Basic_IPv6Address(t *testing.T) {
	p := &parser.ShadowsocksRParser{}
	input := "ssr://WzIwMDE6MGRiODo4NWEzOjAwMDA6MDAwMDo4YTJlOjAzNzA6NzMzNF06NDQzOm9yaWdpbjphZXMtMTkyLWNmYjpwbGFpbjpNVEl6TVRJei8/Z3JvdXA9WkdWbVlYVnNkQSZyZW1hcmtzPVNFRklRUQ"

	expected := proxy.Proxy{
		Type: "ssr",
		Name: "HAHA",
		ShadowSocksR: proxy.ShadowSocksR{
			Server:        "[2001:0db8:85a3:0000:0000:8a2e:0370:7334]",
			Port:          443,
			Cipher:        "aes-192-cfb",
			Password:      "123123",
			ObfsParam:     "",
			Obfs:          "plain",
			Protocol:      "origin",
			ProtocolParam: "",
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestShadowsocksR_Error_InvalidBase64(t *testing.T) {
	p := &parser.ShadowsocksRParser{}
	input := "ssr://invalid_base64"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestShadowsocksR_Error_InvalidProtocol(t *testing.T) {
	p := &parser.ShadowsocksRParser{}
	input := "ss://example.com:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
