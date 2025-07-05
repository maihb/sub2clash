package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestAnytls_Basic_SimpleLink(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@127.0.0.1:8080#Anytls%20Proxy"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "Anytls Proxy",
		Anytls: proxy.Anytls{
			Server:         "127.0.0.1",
			Port:           8080,
			Password:       "password123",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Basic_WithSNI(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@proxy.example.com:443?sni=proxy.example.com#Anytls%20SNI"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "Anytls SNI",
		Anytls: proxy.Anytls{
			Server:         "proxy.example.com",
			Port:           443,
			Password:       "password123",
			SNI:            "proxy.example.com",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Basic_WithInsecure(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@proxy.example.com:443?insecure=1&sni=proxy.example.com#Anytls%20Insecure"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "Anytls Insecure",
		Anytls: proxy.Anytls{
			Server:         "proxy.example.com",
			Port:           443,
			Password:       "password123",
			SNI:            "proxy.example.com",
			SkipCertVerify: true,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Basic_IPv6Address(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@[2001:db8::1]:8080#Anytls%20IPv6"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "Anytls IPv6",
		Anytls: proxy.Anytls{
			Server:         "2001:db8::1",
			Port:           8080,
			Password:       "password123",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Basic_ComplexPassword(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://ComplexPassword!%40%23%24@proxy.example.com:8443?sni=example.com&insecure=1#Anytls%20Full"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "Anytls Full",
		Anytls: proxy.Anytls{
			Server:         "proxy.example.com",
			Port:           8443,
			Password:       "ComplexPassword!@#$",
			SNI:            "example.com",
			SkipCertVerify: true,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Basic_NoPassword(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://@127.0.0.1:8080#No%20Password"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "No Password",
		Anytls: proxy.Anytls{
			Server:         "127.0.0.1",
			Port:           8080,
			Password:       "",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Basic_UsernameOnly(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://username@127.0.0.1:8080#Username%20Only"

	expected := proxy.Proxy{
		Type: "anytls",
		Name: "Username Only",
		Anytls: proxy.Anytls{
			Server:         "127.0.0.1",
			Port:           8080,
			Password:       "username",
			SkipCertVerify: false,
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestAnytls_Error_MissingServer(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestAnytls_Error_MissingPort(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@127.0.0.1"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestAnytls_Error_InvalidPort(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anytls://password123@127.0.0.1:99999"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestAnytls_Error_InvalidProtocol(t *testing.T) {
	p := &parser.AnytlsParser{}
	input := "anyssl://example.com:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
