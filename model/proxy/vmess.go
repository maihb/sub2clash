package proxy

type VmessJson struct {
	V    any    `json:"v"`
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port any    `json:"port"`
	Id   string `json:"id"`
	Aid  any    `json:"aid"`
	Scy  string `json:"scy"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	Tls  string `json:"tls"`
	Sni  string `json:"sni"`
	Alpn string `json:"alpn"`
	Fp   string `json:"fp"`
}

type Vmess struct {
	Type                string         `yaml:"type"`
	Name                string         `yaml:"name"`
	Server              string         `yaml:"server"`
	Port                int            `yaml:"port"`
	UUID                string         `yaml:"uuid"`
	AlterID             int            `yaml:"alterId"`
	Cipher              string         `yaml:"cipher"`
	UDP                 bool           `yaml:"udp,omitempty"`
	Network             string         `yaml:"network,omitempty"`
	TLS                 bool           `yaml:"tls,omitempty"`
	ALPN                []string       `yaml:"alpn,omitempty"`
	SkipCertVerify      bool           `yaml:"skip-cert-verify,omitempty"`
	Fingerprint         string         `yaml:"fingerprint,omitempty"`
	ServerName          string         `yaml:"servername,omitempty"`
	RealityOpts         RealityOptions `yaml:"reality-opts,omitempty"`
	HTTPOpts            HTTPOptions    `yaml:"http-opts,omitempty"`
	HTTP2Opts           HTTP2Options   `yaml:"h2-opts,omitempty"`
	GrpcOpts            GrpcOptions    `yaml:"grpc-opts,omitempty"`
	WSOpts              WSOptions      `yaml:"ws-opts,omitempty"`
	PacketAddr          bool           `yaml:"packet-addr,omitempty"`
	XUDP                bool           `yaml:"xudp,omitempty"`
	PacketEncoding      string         `yaml:"packet-encoding,omitempty"`
	GlobalPadding       bool           `yaml:"global-padding,omitempty"`
	AuthenticatedLength bool           `yaml:"authenticated-length,omitempty"`
	ClientFingerprint   string         `yaml:"client-fingerprint,omitempty"`
}

func ProxyToVmess(p Proxy) Vmess {
	return Vmess{
		Type:                "vmess",
		Name:                p.Name,
		Server:              p.Server,
		Port:                p.Port,
		UUID:                p.UUID,
		AlterID:             p.AlterID,
		Cipher:              p.Cipher,
		UDP:                 p.UDP,
		Network:             p.Network,
		TLS:                 p.TLS,
		ALPN:                p.Alpn,
		SkipCertVerify:      p.SkipCertVerify,
		Fingerprint:         p.Fingerprint,
		ServerName:          p.Servername,
		RealityOpts:         p.RealityOpts,
		HTTPOpts:            p.HTTPOpts,
		HTTP2Opts:           p.HTTP2Opts,
		GrpcOpts:            p.GrpcOpts,
		WSOpts:              p.WSOpts,
		PacketAddr:          p.PacketAddr,
		XUDP:                p.XUDP,
		PacketEncoding:      p.PacketEncoding,
		GlobalPadding:       p.GlobalPadding,
		AuthenticatedLength: p.AuthenticatedLength,
		ClientFingerprint:   p.ClientFingerprint,
	}
}

type VmessMarshaler struct{}

func (m *VmessMarshaler) GetType() string {
	return "vmess"
}

func (m *VmessMarshaler) MarshalProxy(p Proxy) (interface{}, error) {
	return ProxyToVmess(p), nil
}
