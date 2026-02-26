# netbird 信号服务器

这是一个 netbird 信号交换服务器和客户端库，用于在 netbird 对等点之间交换连接信息

## 命令选项

CLI 接受以下选项：

```shell
启动 Netbird 信号服务器守护进程

用法：
  netbird-signal run [flags]

标志：
  -h, --help                        帮助信息
      --letsencrypt-domain string   用于颁发 Let's Encrypt 证书的域名。启用使用 Let's Encrypt 的 TLS。将获取和续订证书，并使用 TLS 运行服务器
      --port int                    服务器监听端口（例如 10000）（默认 10000）
      --ssl-dir string              服务器 ssl 目录位置。*仅对 Let's Encrypt 证书需要。（默认 "/var/lib/netbird/"）
      --cert-file string            SSL 证书的位置。当您已有现有证书且不想自动生成新证书时可以使用。如果指定了 letsencrypt-domain，则此属性无效
      --cert-key string             SSL 证书私钥的位置。当您已有现有证书且不想自动生成新证书时可以使用。如果指定了 letsencrypt-domain，则此属性无效

全局标志：
      --log-file string    设置 Netbird 日志路径。如果指定为 console，则日志将输出到 stdout（默认 "/var/log/netbird/signal.log"）
      --log-level string    （默认 "info"）
```

## 运行信号服务（Docker）

我们将信号服务器打包到 docker 镜像中。您可以从 Docker Hub 拉取镜像并使用以下命令执行：

````shell
docker pull netbirdio/signal:latest
docker run -d --name netbird-signal -p 10000:10000 netbirdio/signal:latest
````

默认日志级别设置为 INFO，如果需要，您可以通过更新 docker 命令来更改它：

````shell
docker run -d --name netbird-signal -p 10000:10000 netbirdio/signal:latest --log-level DEBUG
````

### 使用 TLS 运行（Let's Encrypt）。

通过指定 **--letsencrypt-domain**，守护进程将处理 SSL 证书请求和配置。

在以下示例中，```10000``` 是信号服务的**默认**端口，```443``` 将用于 Let's Encrypt 挑战和 HTTP API。
> 运行容器的服务器必须具有公网 IP（用于 Let's Encrypt 证书挑战）。

将 `<YOUR-DOMAIN>` 替换为您服务器的公网域名（例如 mydomain.com 或子域名 sub.mydomain.com）。

```bash
# 创建卷
docker volume create netbird-signal
# 运行 docker 容器
docker run -d --name netbird-signal \
-p 10000:10000  \
-p 443:443  \
-v netbird-signal:/var/lib/netbird  \
netbirdio/signal:latest \
--letsencrypt-domain <YOUR-DOMAIN>
```

## 指标

信号服务器以 Prometheus 格式公开以下指标：

### 应用程序指标

- **active_peers**：一个 Gauge 指标，跟踪连接到服务器的活动对等点数量。
- **peer_connection_duration_seconds**：一个 Histogram 指标，测量对等点连接持续的时间（秒）。
- **registrations_total**：一个 Counter 指标，计算对等点注册的总次数。
- **deregistrations_total**：一个 Counter 指标，计算对等点注销的总次数。
- **registration_failures_total**：一个 Counter 指标，计算对等点注册失败的总次数。可能的标签：
  - `error`：导致注册失败的错误类型（例如 `missing_id`、`missing_meta`、`failed_header`）。
- **registration_delay_milliseconds**：一个 Histogram 指标，测量注册对等点所需的时间（毫秒）。
- **get_registration_delay_milliseconds**：一个 Histogram 指标，测量获取对等点注册所需的时间（毫秒）。
- **messages_forwarded_total**：一个 Counter 指标，计算对等点之间转发消息的总次数。
- **message_forward_failures_total**：一个 Counter 指标，计算对等点之间消息转发失败的总次数。可能的标签：
  - `type`：失败的类型（例如 `error`、`not_connected`、`not_registered`）。
- **message_forward_latency_milliseconds**：一个 Histogram 指标，测量对等点之间消息转发的延迟（毫秒）。

### 端点

指标以 Prometheus 格式暴露在 `/metrics` 端点上。默认情况下，服务器监听端口 `9090`，因此完整的端点将是：

> http://<server_ip>:9090/metrics

## 用于开发目的：

该项目使用 gRpc 库，并在以下位置的 protobuf 文件中定义服务：
```proto/signalexchange.proto```

要构建项目，您需要执行以下操作。

安装 golang gRpc 工具：

```bash
#!/bin/bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

生成 gRpc 代码：

```bash
#!/bin/bash
protoc -I proto/ proto/signalexchange.proto --go_out=. --go-grpc_out=.
```
