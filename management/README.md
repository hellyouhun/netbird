# netbird 管理服务器
netbird 管理服务器将控制和同步您的 Netbird 账户和网络中的对等点配置。

## 命令选项
CLI 接受 **management** 命令，具有以下选项：
```shell
启动 Netbird 管理服务器

用法：
  netbird-mgmt management [flags]

标志：
      --cert-file string            SSL 证书的位置。当您已有现有证书且不想自动生成新证书时可以使用。如果指定了 letsencrypt-domain，则此属性无效
      --cert-key string             SSL 证书私钥的位置。当您已有现有证书且不想自动生成新证书时可以使用。如果指定了 letsencrypt-domain，则此属性无效
      --datadir string              服务器数据目录位置
  -h, --help                        帮助信息
      --letsencrypt-domain string   用于颁发 Let's Encrypt 证书的域名。启用使用 Let's Encrypt 的 TLS。将获取和续订证书，并使用 TLS 运行服务器
      --port int                    服务器监听端口（默认 33073）

全局标志：
      --config string      Netbird 配置文件位置，用于写入新配置（默认 "/etc/netbird"）
      --log-file string    设置 Netbird 日志路径。如果指定为 console，则日志将输出到 stdout（默认 "/var/log/netbird/management.log"）
      --log-level string    （默认 "info"）
```
## 运行管理服务（Docker）

您可以在 2 种模式下运行服务 - 使用 TLS 或不使用（不推荐）。

### 使用 TLS 运行（Let's Encrypt）。
通过指定 **--letsencrypt-domain**，守护进程将处理 SSL 证书请求和配置。

在以下示例中，```33073``` 是管理服务的**默认**端口，```443``` 将用于 Let's Encrypt 挑战和 HTTP API。
> 运行容器的服务器必须具有公网 IP（用于 Let's Encrypt 证书挑战）。

将 <YOUR-DOMAIN> 替换为您服务器的公网域名（例如 mydomain.com 或子域名 sub.mydomain.com）。

```bash
# 创建卷
docker volume create netbird-mgmt
# 运行 docker 容器
docker run -d --name netbird-management \
-p 33073:33073  \
-p 443:443  \
-v netbird-mgmt:/var/lib/netbird  \
-v ./config.json:/etc/netbird/config.json  \
netbirdio/management:latest \
--letsencrypt-domain <YOUR-DOMAIN>
```
> config.json 的示例可以在 [management.json](../infrastructure_files/management.json.tmpl) 找到

触发 Let's Encrypt 证书生成：
```bash
curl https://<YOUR-DOMAIN>
```

证书将持久化在容器内的 ```datadir/letsencrypt/``` 文件夹中（例如 ```/var/lib/netbird/letsencrypt/```）。

确保 ```datadir``` 映射到主机上的某个文件夹。如果您使用了卷命令，可以运行以下命令来检索挂载点：
```shell
docker volume inspect netbird-mgmt
[
    {
        "CreatedAt": "2021-07-25T20:45:28Z",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/mgmt/_data",
        "Name": "netbird-mgmt",
        "Options": {},
        "Scope": "local"
    }
]
```
容器的后续重启将使用先前生成的证书，因此无需在每次重启时使用 ```curl``` 命令触发证书生成。

### 不使用 TLS 运行。

```bash
# 创建卷
docker volume create netbird-mgmt
# 运行 docker 容器
docker run -d --name netbird-management \
-p 33073:33073  \
-v netbird-mgmt:/var/lib/netbird  \
-v ./config.json:/etc/netbird/config.json  \
netbirdio/management:latest
```
### 调试标签
我们还发布带有调试标签的 docker 镜像，该镜像的日志级别设置为默认值，并且使用 ```gcr.io/distroless/base:debug``` 镜像，可以与 docker exec 一起使用，以便在管理容器中运行一些命令。
```shell
shell $ docker run -d --name netbird-management-debug \
-p 33073:33073  \
-v netbird-mgmt:/var/lib/netbird  \
-v ./config.json:/etc/netbird/config.json  \
netbirdio/management:debug-latest

shell $ docker exec -ti netbird-management-debug /bin/sh
container-shell $ 
```
## 用于开发目的：

安装 golang gRpc 工具：
```bash
#!/bin/bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

生成 gRpc 代码：

```bash
#!/bin/bash
protoc -I proto/ proto/management.proto --go_out=. --go-grpc_out=.
```


