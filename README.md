# sub2clash

将订阅链接转换为 Clash、Clash.Meta 配置  
[预览](https://clash.nite07.com/)

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

- [docker compose](./compose.yml)
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
2. **环境变量**：使用 `SUB2CLASH_` 前缀，例如 `SUB2CLASH_ADDRESS=0.0.0.0:8011`
3. **默认值**：内置默认配置

| 配置项                | 环境变量                        | 说明                                    | 默认值                                                                                               |
| --------------------- | ------------------------------- | --------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| address               | SUB2CLASH_ADDRESS               | 服务监听地址                            | `0.0.0.0:8011`                                                                                       |
| meta_template         | SUB2CLASH_META_TEMPLATE         | 默认 meta 模板 URL                      | `https://raw.githubusercontent.com/bestnite/sub2clash/refs/heads/main/templates/template_meta.yaml`  |
| clash_template        | SUB2CLASH_CLASH_TEMPLATE        | 默认 clash 模板 URL                     | `https://raw.githubusercontent.com/bestnite/sub2clash/refs/heads/main/templates/template_clash.yaml` |
| request_retry_times   | SUB2CLASH_REQUEST_RETRY_TIMES   | 请求重试次数                            | `3`                                                                                                  |
| request_max_file_size | SUB2CLASH_REQUEST_MAX_FILE_SIZE | 请求文件最大大小（byte）                | `1048576`                                                                                            |
| cache_expire          | SUB2CLASH_CACHE_EXPIRE          | 订阅缓存时间（秒）                      | `300`                                                                                                |
| log_level             | SUB2CLASH_LOG_LEVEL             | 日志等级：`debug`,`info`,`warn`,`error` | `info`                                                                                               |
| short_link_length     | SUB2CLASH_SHORT_LINK_LENGTH     | 短链长度                                | `6`                                                                                                  |

#### 配置文件示例

参考示例文件：

- [config.example.yaml](./config.example.yaml) - YAML 格式
- [config.example.json](./config.example.json) - JSON 格式

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
