package proxy

import "fmt"

type HTTPOptions struct {
	Method  string              `yaml:"method,omitempty"`
	Path    []string            `yaml:"path,omitempty"`
	Headers map[string][]string `yaml:"headers,omitempty"`
}

type HTTP2Options struct {
	Host []string `yaml:"host,omitempty"`
	Path string   `yaml:"path,omitempty"`
}

type GrpcOptions struct {
	GrpcServiceName string `yaml:"grpc-service-name,omitempty"`
}

type RealityOptions struct {
	PublicKey string `yaml:"public-key"`
	ShortID   string `yaml:"short-id,omitempty"`
}

type WSOptions struct {
	Path                string            `yaml:"path,omitempty"`
	Headers             map[string]string `yaml:"headers,omitempty"`
	MaxEarlyData        int               `yaml:"max-early-data,omitempty"`
	EarlyDataHeaderName string            `yaml:"early-data-header-name,omitempty"`
}

type SmuxStruct struct {
	Enabled bool `yaml:"enable"`
}

type WireGuardPeerOption struct {
	Server       string   `yaml:"server"`
	Port         int      `yaml:"port"`
	PublicKey    string   `yaml:"public-key,omitempty"`
	PreSharedKey string   `yaml:"pre-shared-key,omitempty"`
	Reserved     []uint8  `yaml:"reserved,omitempty"`
	AllowedIPs   []string `yaml:"allowed-ips,omitempty"`
}

type ECHOptions struct {
	Enable bool   `yaml:"enable,omitempty" obfs:"enable,omitempty"`
	Config string `yaml:"config,omitempty" obfs:"config,omitempty"`
}

type Proxy struct {
	Type    string
	Name    string
	SubName string `yaml:"-"`
	Anytls
	Hysteria
	Hysteria2
	ShadowSocks
	ShadowSocksR
	Trojan
	Vless
	Vmess
	Socks
}

func (p Proxy) MarshalYAML() (any, error) {
	switch p.Type {
	case "anytls":
		return struct {
			Type   string `yaml:"type"`
			Name   string `yaml:"name"`
			Anytls `yaml:",inline"`
		}{
			Type:   p.Type,
			Name:   p.Name,
			Anytls: p.Anytls,
		}, nil
	case "hysteria":
		return struct {
			Type     string `yaml:"type"`
			Name     string `yaml:"name"`
			Hysteria `yaml:",inline"`
		}{
			Type:     p.Type,
			Name:     p.Name,
			Hysteria: p.Hysteria,
		}, nil
	case "hysteria2":
		return struct {
			Type      string `yaml:"type"`
			Name      string `yaml:"name"`
			Hysteria2 `yaml:",inline"`
		}{
			Type:      p.Type,
			Name:      p.Name,
			Hysteria2: p.Hysteria2,
		}, nil
	case "ss":
		return struct {
			Type        string `yaml:"type"`
			Name        string `yaml:"name"`
			ShadowSocks `yaml:",inline"`
		}{
			Type:        p.Type,
			Name:        p.Name,
			ShadowSocks: p.ShadowSocks,
		}, nil
	case "ssr":
		return struct {
			Type         string `yaml:"type"`
			Name         string `yaml:"name"`
			ShadowSocksR `yaml:",inline"`
		}{
			Type:         p.Type,
			Name:         p.Name,
			ShadowSocksR: p.ShadowSocksR,
		}, nil
	case "trojan":
		return struct {
			Type   string `yaml:"type"`
			Name   string `yaml:"name"`
			Trojan `yaml:",inline"`
		}{
			Type:   p.Type,
			Name:   p.Name,
			Trojan: p.Trojan,
		}, nil
	case "vless":
		return struct {
			Type  string `yaml:"type"`
			Name  string `yaml:"name"`
			Vless `yaml:",inline"`
		}{
			Type:  p.Type,
			Name:  p.Name,
			Vless: p.Vless,
		}, nil
	case "vmess":
		return struct {
			Type  string `yaml:"type"`
			Name  string `yaml:"name"`
			Vmess `yaml:",inline"`
		}{
			Type:  p.Type,
			Name:  p.Name,
			Vmess: p.Vmess,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported proxy type: %s", p.Type)
	}
}
