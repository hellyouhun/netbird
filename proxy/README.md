# Netbird 反向代理

NetBird 反向代理是一个单独的服务，可以作为 NetBird 网络中某些资源的公共入口点。
从高层次来看，它的运作方式是：
- 配置的路由从管理服务器传递到代理。
- 对于每个路由，代理会创建到托管资源的 NetBird 对等点的 NetBird 连接。
- 当流量到达代理上为代理资源配置的地址和路径时，NetBird 代理会为该资源启动相关的身份验证方法。
- 成功认证后，代理将流量转发到 NetBird 对等点。

代理支持的身份验证方法包括：
- 无身份验证
- Oauth2/OIDC
- 电子邮件魔法链接
- 简单 PIN 码
- HTTP 基本认证用户名和密码

## 管理连接和身份验证

代理通过 gRPC 连接与管理服务器通信。
代理作为管理服务器的客户端，使用以下 RPC：
- 用于代理服务更新的服务器端流。
- 用于代理日志的客户端流。

要与管理服务器进行身份验证，代理服务器使用机器对机器 OAuth2。
如果您使用的是嵌入式 IdP //TODO：解释如何获取凭据。
否则，在您的 IdP 中为代理服务器创建一个新的机器对机器配置文件，并在代理的环境或标志中设置相关设置（见下文）。

## 用户身份验证

当请求到达代理时，它会查找主机域允许的身份验证方法。
如果没有为主机域注册身份验证方法，则不会应用身份验证（用于完全公开的资源）。
如果为主机域注册了任何身份验证方法，则代理将首先提供身份验证页面，允许用户选择身份验证方法（从允许的方法中）并输入该身份验证方法所需的信息。
如果用户成功通过身份验证，他们的请求将被转发到代理以代理到相关对等点。
成功的身份验证并不能保证请求的成功转发，因为代理后面可能存在故障，例如对等点连接或底层资源问题。

## TLS

由于提供了身份验证，代理对其端点使用 HTTPS，即使底层服务是 HTTP。
证书生成可以通过 ACME（默认情况下使用 Let's Encrypt，但可以使用其他 ACME 提供商）或通过证书文件完成。
当不使用 ACME 时，代理服务器尝试从指定证书目录中的文件 `tls.crt` 和 `tls.key` 加载证书和密钥。
使用 ACME 时，代理服务器将在指定的证书目录中存储生成的证书。


## 认证 UI

认证 UI 是一个位于 `web/` 目录中的 Vite + React 应用程序。它在构建时嵌入到 Go 二进制文件中。

要构建 UI：
```bash
cd web
npm install
npm run build
```

用于热重载的 UI 开发（在 http://localhost:3031 提供服务）：
```bash
npm run dev
```

`web/dist/` 中的构建资源通过 `//go:embed` 嵌入，并由 `web.ServeHTTP` 处理程序提供服务。

## 配置

NetBird 代理部署配置通过标志或环境变量进行，标志优先于环境变量。
以下部署配置可用：

| 标志             | 环境变量                              | 目的                                                                                                                            | 默认值                                            |
|------------------|----------------------------------|------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| `-debug`         | `NB_PROXY_DEBUG_LOGS`            | 启用调试日志                                                                                                               | `false`                                            |
| `-mgmt`          | `NB_PROXY_MANAGEMENT_ADDRESS`    | 代理从中获取配置的管理服务器地址。                                                      | `"https://api.netbird.io:443"`                     |
| `-addr`          | `NB_PROXY_ADDRESS`               | 反向代理将监听的地址。                                                                                 | `":443`                                            |
| `-url`           | `NB_PROXY_URL`                   | 代理将被访问的 URL（端点将被 CNAMEd 到的地方）。如果未设置，这将回退到代理地址。 | `"proxy.netbird.io"`                               |
| `-cert-dir`      | `NB_PROXY_CERTIFICATE_DIRECTORY` | 证书存储的位置。                                                                                      | `"./certs"`                                        |
| `-acme-certs`    | `NB_PROXY_ACME_CERTIFICATES`     | 是否使用 ACME 生成证书。                                                                                      | `false`                                            |
| `-acme-addr`     | `NB_PROXY_ACME_ADDRESS`          | 代理将监听的 HTTP 地址，用于响应 HTTP-01 ACME 挑战                                                    | `":80"`                                            |
| `-acme-dir`      | `NB_PROXY_ACME_DIRECTORY`        | 要使用的 ACME 服务器的目录 URL                                                                                    | `"https://acme-v02.api.letsencrypt.org/directory"` |
| `-oidc-id`       | `NB_PROXY_OIDC_CLIENT_ID`        | OIDC 用户身份验证的 OAuth2 客户端 ID                                                                                  | `"netbird-proxy"`                                  |
| `-oidc-secret`   | `NB_PROXY_OIDC_CLIENT_SECRET`    | OIDC 用户身份验证的 OAuth2 客户端密钥                                                                              | `""`                                               |
| `-oidc-endpoint` | `NB_PROXY_OIDC_ENDPOINT`         | OIDC 用户身份验证的 OAuth2 提供商端点                                                                          | `"https://api.netbird.io/oauth2"`                  |
| `-oidc-scopes`   | `NB_PROXY_OIDC_SCOPES`           | OIDC 用户身份验证的 OAuth2 范围，逗号分隔                                                                    | `"openid,profile,email"`                           |
