package proxy

type ShadowSocks struct {
	Server            string         `yaml:"server"`
	Port              int            `yaml:"port"`
	Password          string         `yaml:"password"`
	Cipher            string         `yaml:"cipher"`
	UDP               bool           `yaml:"udp,omitempty"`
	Plugin            string         `yaml:"plugin,omitempty"`
	PluginOpts        map[string]any `yaml:"plugin-opts,omitempty"`
	UDPOverTCP        bool           `yaml:"udp-over-tcp,omitempty"`
	UDPOverTCPVersion int            `yaml:"udp-over-tcp-version,omitempty"`
	ClientFingerprint string         `yaml:"client-fingerprint,omitempty"`
}
