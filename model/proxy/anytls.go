package proxy

type Anytls struct {
	Server                   string   `yaml:"server"`
	Port                     int      `yaml:"port"`
	Password                 string   `yaml:"password,omitempty"`
	Alpn                     []string `yaml:"alpn,omitempty"`
	Sni                      string   `yaml:"sni,omitempty"`
	ClientFingerprint        string   `yaml:"client-fingerprint,omitempty"`
	SkipCertVerify           bool     `yaml:"skip-cert-verify,omitempty"`
	Fingerprint              string   `yaml:"fingerprint,omitempty"`
	UDP                      bool     `yaml:"udp,omitempty"`
	IdleSessionCheckInterval int      `yaml:"idle-session-check-interval,omitempty"`
	IdleSessionTimeout       int      `yaml:"idle-session-timeout,omitempty"`
	MinIdleSession           int      `yaml:"min-idle-session,omitempty"`
}
