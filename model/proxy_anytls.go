package model

type Anytls struct {
	Type                     string   `yaml:"type"`
	Name                     string   `yaml:"name"`
	Server                   string   `yaml:"server"`
	Port                     int      `yaml:"port"`
	Password                 string   `yaml:"password,omitempty"`
	Alpn                     []string `yaml:"alpn,omitempty"`
	SNI                      string   `yaml:"sni,omitempty"`
	ClientFingerprint        string   `yaml:"client-fingerprint,omitempty"`
	SkipCertVerify           bool     `yaml:"skip-cert-verify,omitempty"`
	Fingerprint              string   `yaml:"fingerprint,omitempty"`
	UDP                      bool     `yaml:"udp,omitempty"`
	IdleSessionCheckInterval int      `yaml:"idle-session-check-interval,omitempty"`
	IdleSessionTimeout       int      `yaml:"idle-session-timeout,omitempty"`
	MinIdleSession           int      `yaml:"min-idle-session,omitempty"`
}

func ProxyToAnytls(p Proxy) Anytls {
	return Anytls{
		Type:                     "anytls",
		Name:                     p.Name,
		Server:                   p.Server,
		Port:                     p.Port,
		Password:                 p.Password,
		Alpn:                     p.Alpn,
		SNI:                      p.Sni,
		ClientFingerprint:        p.ClientFingerprint,
		SkipCertVerify:           p.SkipCertVerify,
		Fingerprint:              p.Fingerprint,
		UDP:                      p.UDP,
		IdleSessionCheckInterval: p.IdleSessionCheckInterval,
		IdleSessionTimeout:       p.IdleSessionTimeout,
		MinIdleSession:           p.MinIdleSession,
	}
}
