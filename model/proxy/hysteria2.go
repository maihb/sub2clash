package proxy

type Hysteria2 struct {
	Server         string   `yaml:"server"`
	Port           int      `yaml:"port"`
	Up             string   `yaml:"up,omitempty"`
	Down           string   `yaml:"down,omitempty"`
	Password       string   `yaml:"password,omitempty"`
	Obfs           string   `yaml:"obfs,omitempty"`
	ObfsPassword   string   `yaml:"obfs-password,omitempty"`
	SNI            string   `yaml:"sni,omitempty"`
	SkipCertVerify bool     `yaml:"skip-cert-verify,omitempty"`
	Fingerprint    string   `yaml:"fingerprint,omitempty"`
	ALPN           []string `yaml:"alpn,omitempty"`
	CustomCA       string   `yaml:"ca,omitempty"`
	CustomCAString string   `yaml:"ca-str,omitempty"`
	CWND           int      `yaml:"cwnd,omitempty"`
	ObfsParam      string   `yaml:"obfs-param,omitempty"`
	Sni            string   `yaml:"sni,omitempty"`
	TLS            bool     `yaml:"tls,omitempty"`
	Network        string   `yaml:"network,omitempty"`
}
