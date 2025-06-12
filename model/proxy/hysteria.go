package proxy

type Hysteria struct {
	Server              string   `yaml:"server"`
	Port                int      `yaml:"port,omitempty"`
	Ports               string   `yaml:"ports,omitempty"`
	Protocol            string   `yaml:"protocol,omitempty"`
	ObfsProtocol        string   `yaml:"obfs-protocol,omitempty"`
	Up                  string   `yaml:"up"`
	UpSpeed             int      `yaml:"up-speed,omitempty"`
	Down                string   `yaml:"down"`
	DownSpeed           int      `yaml:"down-speed,omitempty"`
	Auth                string   `yaml:"auth,omitempty"`
	AuthStringOLD       string   `yaml:"auth_str,omitempty"`
	AuthString          string   `yaml:"auth-str,omitempty"`
	Obfs                string   `yaml:"obfs,omitempty"`
	SNI                 string   `yaml:"sni,omitempty"`
	SkipCertVerify      bool     `yaml:"skip-cert-verify,omitempty"`
	Fingerprint         string   `yaml:"fingerprint,omitempty"`
	Alpn                []string `yaml:"alpn,omitempty"`
	CustomCA            string   `yaml:"ca,omitempty"`
	CustomCAString      string   `yaml:"ca-str,omitempty"`
	ReceiveWindowConn   int      `yaml:"recv-window-conn,omitempty"`
	ReceiveWindow       int      `yaml:"recv-window,omitempty"`
	DisableMTUDiscovery bool     `yaml:"disable-mtu-discovery,omitempty"`
	FastOpen            bool     `yaml:"fast-open,omitempty"`
	HopInterval         int      `yaml:"hop-interval,omitempty"`
	AllowInsecure       bool     `yaml:"allow-insecure,omitempty"`
}
