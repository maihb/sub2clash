package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestVless_Basic_SimpleLink(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:8080#VLESS%20Proxy"

	expected := proxy.Proxy{
		Type: "vless",
		Name: "VLESS Proxy",
		Vless: proxy.Vless{
			Server: "127.0.0.1",
			Port:   8080,
			UUID:   "b831b0c4-33b7-4873-9834-28d66d87d4ce",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVless_Basic_WithTLS(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:443?security=tls&sni=example.com&alpn=h2,http/1.1#VLESS%20TLS"

	expected := proxy.Proxy{
		Type: "vless",
		Name: "VLESS TLS",
		Vless: proxy.Vless{
			Server:     "127.0.0.1",
			Port:       443,
			UUID:       "b831b0c4-33b7-4873-9834-28d66d87d4ce",
			TLS:        true,
			ALPN:       []string{"h2", "http/1.1"},
			ServerName: "example.com",
		},
	}

	result, err := p.Parse(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVless_Basic_WithReality(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:443?security=reality&sni=example.com&pbk=publickey123&sid=shortid123&fp=chrome#VLESS%20Reality"

	expected := proxy.Proxy{
		Type: "vless",
		Name: "VLESS Reality",
		Vless: proxy.Vless{
			Server:     "127.0.0.1",
			Port:       443,
			UUID:       "b831b0c4-33b7-4873-9834-28d66d87d4ce",
			TLS:        true,
			ServerName: "example.com",
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

func TestVless_Basic_WithWebSocket(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:443?type=ws&path=/ws&host=example.com#VLESS%20WS"

	expected := proxy.Proxy{
		Type: "vless",
		Name: "VLESS WS",
		Vless: proxy.Vless{
			Server:  "127.0.0.1",
			Port:    443,
			UUID:    "b831b0c4-33b7-4873-9834-28d66d87d4ce",
			Network: "ws",
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

func TestVless_Basic_WithGrpc(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:443?type=grpc&serviceName=grpc_service#VLESS%20gRPC"

	expected := proxy.Proxy{
		Type: "vless",
		Name: "VLESS gRPC",
		Vless: proxy.Vless{
			Server:  "127.0.0.1",
			Port:    443,
			UUID:    "b831b0c4-33b7-4873-9834-28d66d87d4ce",
			Network: "grpc",
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

func TestVless_Basic_WithHTTP(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:443?type=http&path=/path1,/path2&host=host1.com,host2.com#VLESS%20HTTP"

	expected := proxy.Proxy{
		Type: "vless",
		Name: "VLESS HTTP",
		Vless: proxy.Vless{
			Server:  "127.0.0.1",
			Port:    443,
			UUID:    "b831b0c4-33b7-4873-9834-28d66d87d4ce",
			Network: "http",
			HTTPOpts: proxy.HTTPOptions{
				Path: []string{"/path1", "/path2"},
				Headers: map[string][]string{
					"host": {"host1.com", "host2.com"},
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

func TestVless_Error_MissingServer(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@:8080"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestVless_Error_MissingPort(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestVless_Error_InvalidPort(t *testing.T) {
	p := &parser.VlessParser{}
	input := "vless://b831b0c4-33b7-4873-9834-28d66d87d4ce@127.0.0.1:99999"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestVless_Error_InvalidProtocol(t *testing.T) {
	p := &parser.VlessParser{}
	input := "ss://example.com:8080"

	_, err := p.Parse(input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
