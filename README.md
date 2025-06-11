# sub2clash

将订阅链接转换为 Clash、Clash.Meta 配置  
[预览](https://www.nite07.com/sub)

## 特性

- 开箱即用的规则、策略组配置
- 自动根据节点名称按国家划分策略组
- 多订阅合并
- 自定义 Rule Provider、Rule
- 支持多种协议
  - Shadowsocks
  - ShadowsocksR
  - Vmess
  - Vless （Clash.Meta）
  - Trojan
  - Hysteria （Clash.Meta）
  - Hysteria2 （Clash.Meta）
  - Socks5
  - Anytls （Clash.Meta）

## 使用

### 部署

- [docker compose](./docker-compose.yml)
- 运行[二进制文件](https://github.com/bestnite/sub2clash/releases/latest)

### 配置

支持多种配置方式，按优先级排序：

1. **配置文件**：支持多种格式（YAML、JSON、TOML、INI），按以下优先级搜索：
   - `config.yaml` / `config.yml`
   - `config.json`
   - `config.toml`
   - `config.ini`
   - `sub2clash.yaml` / `sub2clash.yml`
   - `sub2clash.json`
   - `sub2clash.toml`
   - `sub2clash.ini`
2. **环境变量**：使用 `SUB2CLASH_` 前缀，例如 `SUB2CLASH_PORT=8080`
3. **默认值**：内置默认配置

| 配置项                | 环境变量                        | 说明                                    | 默认值         |
| --------------------- | ------------------------------- | --------------------------------------- | -------------- |
| port                  | SUB2CLASH_PORT                  | 服务端口                                | `8011`         |
| meta_template         | SUB2CLASH_META_TEMPLATE         | 默认 meta 模板 URL                      | [默认模板链接] |
| clash_template        | SUB2CLASH_CLASH_TEMPLATE        | 默认 clash 模板 URL                     | [默认模板链接] |
| request_retry_times   | SUB2CLASH_REQUEST_RETRY_TIMES   | 请求重试次数                            | `3`            |
| request_max_file_size | SUB2CLASH_REQUEST_MAX_FILE_SIZE | 请求文件最大大小（byte）                | `1048576`      |
| cache_expire          | SUB2CLASH_CACHE_EXPIRE          | 订阅缓存时间（秒）                      | `300`          |
| log_level             | SUB2CLASH_LOG_LEVEL             | 日志等级：`debug`,`info`,`warn`,`error` | `info`         |
| short_link_length     | SUB2CLASH_SHORT_LINK_LENGTH     | 短链长度                                | `6`            |

#### 配置文件示例

参考示例文件：

- [config.example.yaml](./config.example.yaml) - YAML 格式
- [config.example.json](./config.example.json) - JSON 格式
- [config.example.toml](./config.example.toml) - TOML 格式

**YAML 格式** (`config.yaml`):

```yaml
port: 8011
meta_template: "https://example.com/meta.yaml"
clash_template: "https://example.com/clash.yaml"
request_retry_times: 5
request_max_file_size: 2097152 # 2MB
cache_expire: 600 # 10 minutes
log_level: "debug"
short_link_length: 8
```

**JSON 格式** (`config.json`):

```json
{
  "port": 8011,
  "meta_template": "https://example.com/meta.yaml",
  "clash_template": "https://example.com/clash.yaml",
  "request_retry_times": 5,
  "request_max_file_size": 2097152,
  "cache_expire": 600,
  "log_level": "debug",
  "short_link_length": 8
}
```

**TOML 格式** (`config.toml`):

```toml
port = 8011
meta_template = "https://example.com/meta.yaml"
clash_template = "https://example.com/clash.yaml"
request_retry_times = 5
request_max_file_size = 2097152
cache_expire = 600
log_level = "debug"
short_link_length = 8
```

### API

[API 文档](./API.md)

### 模板

可以通过变量自定义模板中的策略组代理节点  
具体参考下方默认模板

- `<all>` 为添加所有节点
- `<countries>` 为添加所有国家策略组
- `<地区二位字母代码>` 为添加指定地区所有节点，例如 `<hk>` 将添加所有香港节点

#### 默认模板

- [Clash](./templates/template_clash.yaml)
- [Clash.Meta](./templates/template_meta.yaml)
