package test

import (
	"testing"

	"github.com/maihb/sub2clash/model/proxy"
	"github.com/maihb/sub2clash/parser"
)

func TestVmess_Basic_SimpleLink(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJhZGQiOiIxMjcuMC4wLjEiLCJhaWQiOiIwIiwiaWQiOiIxMjM0NTY3OC05MDEyLTM0NTYtNzg5MC0xMjM0NTY3ODkwMTIiLCJuZXQiOiJ3cyIsInBvcnQiOiI0NDMiLCJwcyI6IkhBSEEiLCJ0bHMiOiJ0bHMiLCJ0eXBlIjoibm9uZSIsInYiOiIyIn0="

	expected := proxy.Proxy{
		Type: "vmess",
		Name: "HAHA",
		Vmess: proxy.Vmess{
			UUID:    "12345678-9012-3456-7890-123456789012",
			AlterID: 0,
			Cipher:  "auto",
			Server:  "127.0.0.1",
			Port:    443,
			TLS:     true,
			Network: "ws",
			WSOpts: proxy.WSOptions{
				Path: "/",
				Headers: map[string]string{
					"Host": "127.0.0.1",
				},
			},
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVmess_Basic_WithPath(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJhZGQiOiIxMjcuMC4wLjEiLCJhaWQiOiIwIiwiaWQiOiIxMjM0NTY3OC05MDEyLTM0NTYtNzg5MC0xMjM0NTY3ODkwMTIiLCJuZXQiOiJ3cyIsInBhdGgiOiIvd3MiLCJwb3J0IjoiNDQzIiwicHMiOiJIQUNLIiwidGxzIjoidGxzIiwidHlwZSI6Im5vbmUiLCJ2IjoiMiJ9"

	expected := proxy.Proxy{
		Type: "vmess",
		Name: "HACK",
		Vmess: proxy.Vmess{
			UUID:    "12345678-9012-3456-7890-123456789012",
			AlterID: 0,
			Cipher:  "auto",
			Server:  "127.0.0.1",
			Port:    443,
			TLS:     true,
			Network: "ws",
			WSOpts: proxy.WSOptions{
				Path: "/ws",
				Headers: map[string]string{
					"Host": "127.0.0.1",
				},
			},
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVmess_Basic_WithHost(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJhZGQiOiIxMjcuMC4wLjEiLCJhaWQiOiIwIiwiaG9zdCI6ImV4YW1wbGUuY29tIiwiaWQiOiIxMjM0NTY3OC05MDEyLTM0NTYtNzg5MC0xMjM0NTY3ODkwMTIiLCJuZXQiOiJ3cyIsInBvcnQiOiI0NDMiLCJwcyI6IkhBSEEiLCJ0bHMiOiJ0bHMiLCJ0eXBlIjoibm9uZSIsInYiOiIyIn0="

	expected := proxy.Proxy{
		Type: "vmess",
		Name: "HAHA",
		Vmess: proxy.Vmess{
			UUID:    "12345678-9012-3456-7890-123456789012",
			AlterID: 0,
			Cipher:  "auto",
			Server:  "127.0.0.1",
			Port:    443,
			TLS:     true,
			Network: "ws",
			WSOpts: proxy.WSOptions{
				Path: "/",
				Headers: map[string]string{
					"Host": "example.com",
				},
			},
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVmess_Basic_WithSNI(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJhZGQiOiIxMjcuMC4wLjEiLCJhaWQiOiIwIiwiaWQiOiIxMjM0NTY3OC05MDEyLTM0NTYtNzg5MC0xMjM0NTY3ODkwMTIiLCJuZXQiOiJ3cyIsInBvcnQiOiI0NDMiLCJwcyI6IkhBSEEiLCJzbmkiOiJleGFtcGxlLmNvbSIsInRscyI6InRscyIsInR5cGUiOiJub25lIiwidiI6IjIifQ=="

	expected := proxy.Proxy{
		Type: "vmess",
		Name: "HAHA",
		Vmess: proxy.Vmess{
			UUID:       "12345678-9012-3456-7890-123456789012",
			AlterID:    0,
			Cipher:     "auto",
			Server:     "127.0.0.1",
			Port:       443,
			TLS:        true,
			Network:    "ws",
			ServerName: "example.com",
			WSOpts: proxy.WSOptions{
				Path: "/",
				Headers: map[string]string{
					"Host": "127.0.0.1",
				},
			},
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVmess_Basic_WithAlterID(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJhZGQiOiIxMjcuMC4wLjEiLCJhaWQiOiIxIiwiaWQiOiIxMjM0NTY3OC05MDEyLTM0NTYtNzg5MC0xMjM0NTY3ODkwMTIiLCJuZXQiOiJ3cyIsInBvcnQiOiI0NDMiLCJwcyI6IkhBSEEiLCJ0bHMiOiJ0bHMiLCJ0eXBlIjoibm9uZSIsInYiOiIyIn0="

	expected := proxy.Proxy{
		Type: "vmess",
		Name: "HAHA",
		Vmess: proxy.Vmess{
			UUID:    "12345678-9012-3456-7890-123456789012",
			AlterID: 1,
			Cipher:  "auto",
			Server:  "127.0.0.1",
			Port:    443,
			TLS:     true,
			Network: "ws",
			WSOpts: proxy.WSOptions{
				Path: "/",
				Headers: map[string]string{
					"Host": "127.0.0.1",
				},
			},
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVmess_Basic_GRPC(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJhZGQiOiIxMjcuMC4wLjEiLCJhaWQiOiIwIiwiaWQiOiIxMjM0NTY3OC05MDEyLTM0NTYtNzg5MC0xMjM0NTY3ODkwMTIiLCJuZXQiOiJncnBjIiwicG9ydCI6IjQ0MyIsInBzIjoiSEFIQSIsInRscyI6InRscyIsInR5cGUiOiJub25lIiwidiI6IjIifQ=="

	expected := proxy.Proxy{
		Type: "vmess",
		Name: "HAHA",
		Vmess: proxy.Vmess{
			UUID:     "12345678-9012-3456-7890-123456789012",
			AlterID:  0,
			Cipher:   "auto",
			Server:   "127.0.0.1",
			Port:     443,
			TLS:      true,
			Network:  "grpc",
			GrpcOpts: proxy.GrpcOptions{},
		},
	}

	result, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	validateResult(t, expected, result)
}

func TestVmess_Error_InvalidBase64(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://invalid_base64"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestVmess_Error_InvalidJSON(t *testing.T) {
	p := &parser.VmessParser{}
	input := "vmess://eyJpbnZhbGlkIjoianNvbn0="

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestVmess_Error_InvalidProtocol(t *testing.T) {
	p := &parser.VmessParser{}
	input := "ss://example.com:8080"

	_, err := p.Parse(parser.ParseConfig{UseUDP: false}, input)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
