package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Address            string `mapstructure:"address"`
	MetaTemplate       string `mapstructure:"meta_template"`
	ClashTemplate      string `mapstructure:"clash_template"`
	RequestRetryTimes  int    `mapstructure:"request_retry_times"`
	RequestMaxFileSize int64  `mapstructure:"request_max_file_size"`
	CacheExpire        int64  `mapstructure:"cache_expire"`
	LogLevel           string `mapstructure:"log_level"`
	ShortLinkLength    int    `mapstructure:"short_link_length"`
}

var GlobalConfig *Config
var Dev string

func LoadConfig() error {
	v := viper.New()

	// 添加配置文件搜索路径
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("/etc/sub2clash/")

	// 设置默认值
	setDefaults(v)

	// 设置环境变量前缀和自动绑定
	v.SetEnvPrefix("SUB2CLASH")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 尝试按优先级加载不同格式的配置文件
	configLoaded := false
	configNames := []string{"config", "sub2clash"}
	configExts := []string{"yaml", "yml", "json"}

	for _, name := range configNames {
		for _, ext := range configExts {
			v.SetConfigName(name)
			v.SetConfigType(ext)
			if err := v.ReadInConfig(); err == nil {
				configLoaded = true
				break
			}
		}
		if configLoaded {
			break
		}
	}

	// 将配置解析到结构体
	GlobalConfig = &Config{}
	if err := v.Unmarshal(GlobalConfig); err != nil {
		return err
	}

	return nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("address", "0.0.0.0:8011")
	v.SetDefault("meta_template", "https://raw.githubusercontent.com/bestnite/sub2clash/refs/heads/main/templates/template_meta.yaml")
	v.SetDefault("clash_template", "https://raw.githubusercontent.com/bestnite/sub2clash/refs/heads/main/templates/template_clash.yaml")
	v.SetDefault("request_retry_times", 3)
	v.SetDefault("request_max_file_size", 1024*1024*1)
	v.SetDefault("cache_expire", 60*5)
	v.SetDefault("log_level", "info")
	v.SetDefault("short_link_length", 6)
}
