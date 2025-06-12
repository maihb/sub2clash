package test

import (
	"testing"

	"github.com/bestnite/sub2clash/model/proxy"
	"github.com/bestnite/sub2clash/parser"
)

func TestSocksParser_Parse(t *testing.T) {
	p := &parser.SocksParser{}

	tests := []struct {
		name        string
		input       string
		expected    proxy.Proxy
		expectError bool
		errorType   string
	}{
		{
			name:  "无认证SOCKS5代理 - IPv4",
			input: "socks://127.0.0.1:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "127.0.0.1:1080",
				Socks: proxy.Socks{
					Server:   "127.0.0.1",
					Port:     1080,
					UserName: "",
					Password: "",
				},
			},
			expectError: false,
		},
		{
			name:  "无认证SOCKS5代理 - IPv6",
			input: "socks://[2001:db8::1]:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "2001:db8::1:1080",
				Socks: proxy.Socks{
					Server:   "2001:db8::1",
					Port:     1080,
					UserName: "",
					Password: "",
				},
			},
			expectError: false,
		},
		{
			name:  "明文用户名密码认证 - IPv4",
			input: "socks://user:pass@127.0.0.1:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "127.0.0.1:1080",
				Socks: proxy.Socks{
					Server:   "127.0.0.1",
					Port:     1080,
					UserName: "user",
					Password: "pass",
				},
			},
			expectError: false,
		},
		{
			name:  "明文用户名密码认证 - IPv6",
			input: "socks://user:pass@[2001:db8::1]:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "2001:db8::1:1080",
				Socks: proxy.Socks{
					Server:   "2001:db8::1",
					Port:     1080,
					UserName: "user",
					Password: "pass",
				},
			},
			expectError: false,
		},
		{
			name:  "Base64编码的用户名密码",
			input: "socks://dXNlcm5hbWU6cGFzc3dvcmQ=@127.0.0.1:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "127.0.0.1:1080",
				Socks: proxy.Socks{
					Server:   "127.0.0.1",
					Port:     1080,
					UserName: "username",
					Password: "password",
				},
			},
			expectError: false,
		},
		{
			name:  "带Fragment的URL",
			input: "socks://username:password@proxy.example.com:1080#MyProxy",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "MyProxy",
				Socks: proxy.Socks{
					Server:   "proxy.example.com",
					Port:     1080,
					UserName: "username",
					Password: "password",
				},
			},
			expectError: false,
		},
		{
			name:  "只有用户名没有密码",
			input: "socks://username@127.0.0.1:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "127.0.0.1:1080",
				Socks: proxy.Socks{
					Server:   "127.0.0.1",
					Port:     1080,
					UserName: "username",
					Password: "",
				},
			},
			expectError: false,
		},
		{
			name:  "Base64编码只有用户名",
			input: "socks://dXNlcm5hbWU=@127.0.0.1:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "127.0.0.1:1080",
				Socks: proxy.Socks{
					Server:   "127.0.0.1",
					Port:     1080,
					UserName: "username",
					Password: "",
				},
			},
			expectError: false,
		},
		{
			name:  "域名代理",
			input: "socks://proxy.example.com:8080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "proxy.example.com:8080",
				Socks: proxy.Socks{
					Server:   "proxy.example.com",
					Port:     8080,
					UserName: "",
					Password: "",
				},
			},
			expectError: false,
		},
		{
			name:  "特殊字符用户名密码",
			input: "socks://user%40domain:p%40ss%21@127.0.0.1:1080",
			expected: proxy.Proxy{
				Type: "socks5",
				Name: "127.0.0.1:1080",
				Socks: proxy.Socks{
					Server:   "127.0.0.1",
					Port:     1080,
					UserName: "user@domain",
					Password: "p@ss!",
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := p.Parse(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("期望出现错误，但没有错误")
					return
				}

				if parseErr, ok := err.(*parser.ParseError); ok {
					if tt.errorType != "" && string(parseErr.Type) != tt.errorType {
						t.Errorf("错误类型不匹配: 期望 %s, 实际 %s", tt.errorType, string(parseErr.Type))
					}
				}
				return
			}

			if err != nil {
				t.Errorf("意外的错误: %v", err)
				return
			}

			// 验证解析结果
			if result.Type != tt.expected.Type {
				t.Errorf("Type不匹配: 期望 %s, 实际 %s", tt.expected.Type, result.Type)
			}

			if result.Name != tt.expected.Name {
				t.Errorf("Name不匹配: 期望 %s, 实际 %s", tt.expected.Name, result.Name)
			}

			if result.Socks.Server != tt.expected.Socks.Server {
				t.Errorf("Server不匹配: 期望 %s, 实际 %s", tt.expected.Socks.Server, result.Socks.Server)
			}

			if result.Socks.Port != tt.expected.Socks.Port {
				t.Errorf("Port不匹配: 期望 %d, 实际 %d", tt.expected.Socks.Port, result.Socks.Port)
			}

			if result.Socks.UserName != tt.expected.Socks.UserName {
				t.Errorf("UserName不匹配: 期望 %s, 实际 %s", tt.expected.Socks.UserName, result.Socks.UserName)
			}

			if result.Socks.Password != tt.expected.Socks.Password {
				t.Errorf("Password不匹配: 期望 %s, 实际 %s", tt.expected.Socks.Password, result.Socks.Password)
			}
		})
	}
}
