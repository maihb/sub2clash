package proxy

type Trojan struct {
	Server            string         `yaml:"server"`
	Port              int            `yaml:"port"`
	Password          string         `yaml:"password"`
	Alpn              []string       `yaml:"alpn,omitempty"`
	Sni               string         `yaml:"sni,omitempty"`
	SkipCertVerify    bool           `yaml:"skip-cert-verify,omitempty"`
	Fingerprint       string         `yaml:"fingerprint,omitempty"`
	UDP               bool           `yaml:"udp,omitempty"`
	Network           string         `yaml:"network,omitempty"`
	RealityOpts       RealityOptions `yaml:"reality-opts,omitempty"`
	GrpcOpts          GrpcOptions    `yaml:"grpc-opts,omitempty"`
	WSOpts            WSOptions      `yaml:"ws-opts,omitempty"`
	ClientFingerprint string         `yaml:"client-fingerprint,omitempty"`
	TLS               bool           `yaml:"tls,omitempty"`
}
