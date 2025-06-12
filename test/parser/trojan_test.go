package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestTrojan_Basic_SimpleLink(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1:443#Trojan%20Proxy"

	expected := proxy.Proxy{
		Type: "trojan",
		Name: "Trojan Proxy",
		Trojan: proxy.Trojan{
			Server:   "127.0.0.1",
			Port:     443,
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

func TestTrojan_Basic_WithTLS(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1:443?security=tls&sni=example.com&alpn=h2,http/1.1#Trojan%20TLS"

	expected := proxy.Proxy{
		Type: "trojan",
		Name: "Trojan TLS",
		Trojan: proxy.Trojan{
			Server:   "127.0.0.1",
			Port:     443,
			Password: "password",
			ALPN:     []string{"h2", "http/1.1"},
			SNI:      "example.com",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTrojan_Basic_WithReality(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1:443?security=reality&sni=example.com&pbk=publickey123&sid=shortid123&fp=chrome#Trojan%20Reality"

	expected := proxy.Proxy{
		Type: "trojan",
		Name: "Trojan Reality",
		Trojan: proxy.Trojan{
			Server:   "127.0.0.1",
			Port:     443,
			Password: "password",
			SNI:      "example.com",
			RealityOpts: proxy.RealityOptions{
				PublicKey: "publickey123",
				ShortID:   "shortid123",
			},
			Fingerprint: "chrome",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTrojan_Basic_WithWebSocket(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1:443?type=ws&path=/ws&host=example.com#Trojan%20WS"

	expected := proxy.Proxy{
		Type: "trojan",
		Name: "Trojan WS",
		Trojan: proxy.Trojan{
			Server:   "127.0.0.1",
			Port:     443,
			Password: "password",
			Network:  "ws",
			WSOpts: proxy.WSOptions{
				Path: "/ws",
				Headers: map[string]string{
					"Host": "example.com",
				},
			},
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTrojan_Basic_WithGrpc(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1:443?type=grpc&serviceName=grpc_service#Trojan%20gRPC"

	expected := proxy.Proxy{
		Type: "trojan",
		Name: "Trojan gRPC",
		Trojan: proxy.Trojan{
			Server:   "127.0.0.1",
			Port:     443,
			Password: "password",
			Network:  "grpc",
			GrpcOpts: proxy.GrpcOptions{
				GrpcServiceName: "grpc_service",
			},
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestTrojan_Error_MissingServer(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@:443"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestTrojan_Error_MissingPort(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestTrojan_Error_InvalidPort(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "trojan://password@127.0.0.1:99999"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestTrojan_Error_InvalidProtocol(t *testing.T) {
	p := &parser.TrojanParser{}
	input := "ss://example.com:8080"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
