# 为 NetBird 做贡献

感谢您有兴趣为 NetBird 做出贡献。

您可以通过多种方式做出贡献：
- 报告问题
- 更新文档
- 在 Slack 或 Reddit 上分享使用案例
- 修复错误或增强功能

如果您尚未加入，请[点击这里](https://join.slack.com/t/netbirdio/shared_invite/zt-vrahf41g-ik1v7fV8du6t0RwxSrJ96A)加入我们的 Slack 工作区，我们很乐意讨论需要社区贡献的主题以及现有功能的增强。

## 目录

- [为 NetBird 做贡献](#为-netbird-做贡献)
    - [目录](#目录)
    - [行为准则](#行为准则)
    - [目录结构](#目录结构)
    - [开发环境设置](#开发环境设置)
        - [要求](#要求)
        - [本地 NetBird 设置](#本地-netbird-设置)
        - [开发容器支持](#开发容器支持)
        - [构建和启动](#构建和启动)
        - [测试套件](#测试套件)
    - [提交 PR 前的检查清单](#提交-pr-前的检查清单)
    - [其他项目仓库](#其他项目仓库)
    - [贡献者许可协议](#贡献者许可协议)

## 行为准则

本项目及其所有参与者均受[行为准则](CODE_OF_CONDUCT.md)的约束。
通过参与，您应该遵守此准则。请将不当行为报告至 community@netbird.io。

## 目录结构

NetBird 项目 monorepo 的组织方式是将大部分独立的依赖代码保留在其各自的目录中，除了一些辅助或共享包之外。

最重要的目录包括：

- [/.github](/.github) - Github 操作工作流文件和问题模板
- [/client](/client) - NetBird 代理代码
- [/client/cmd](/client/cmd) - NetBird 代理 CLI 代码
- [/client/internal](/client/internal) - NetBird 代理业务逻辑代码
- [/client/proto](/client/proto) - NetBird 代理守护进程 GRPC 原型文件
- [/client/server](/client/server) - 用于后台执行的 NetBird 代理守护进程代码
- [/client/ui](/client/ui) - NetBird 代理 UI 代码
- [/encryption](/encryption) - 包含代理通信的主要加密代码
- [/iface](/iface) - Wireguard® 接口代码
- [/infrastructure_files](/infrastructure_files) - 包含 docker 和模板脚本的入门文件
- [/management](/management) - 管理服务代码
- [/management/client](/management/client) - 管理服务客户端代码，由代理代码导入
- [/management/proto](/management/proto) - 管理服务 GRPC 原型文件
- [/management/server](/management/server) - 管理服务服务器代码
- [/management/server/http](/management/server/http) - 管理服务 REST API 代码
- [/management/server/idp](/management/server/idp) - 管理服务 IDP 管理代码
- [/release_files](/release_files) - 进入发布包的文件
- [/signal](/signal) - 信号服务代码
- [/signal/client](/signal/client) - 信号服务客户端代码，由代理代码导入
- [/signal/peer](/signal/peer) - 信号服务对等消息逻辑
- [/signal/proto](/signal/proto) - 信号服务 GRPC 原型文件
- [/signal/server](/signal/server) - 信号服务服务器代码


## 开发环境设置

如果您想为错误修复做出贡献或改进现有功能，您必须确保安装了所有必需的依赖项。以下是如何完成此操作的简短指南。

### 要求

#### Go 1.21

请遵循 https://go.dev/ 的安装指南

#### UI 客户端 - Fyne 工具包

我们在 UI 客户端中使用 fyne 工具包。您可以按照其要求指南安装所有依赖项：https://developer.fyne.io/started/#prerequisites

#### gRPC
您可以按照快速入门指南 https://grpc.io/docs/languages/go/quickstart/#prerequisites 的说明操作，然后运行位于每个 `proto` 目录中的 `generate.sh` 文件以生成更改。
> **重要**：我们非常欢迎能够改进客户端守护进程协议的贡献。对于信号和管理协议，请通过 Slack 或 Github 问题与我们联系。

#### Docker

请遵循 https://docs.docker.com/get-docker/ 的安装指南

#### Goreleaser 和 golangci-lint

我们在 Github 操作工作流中使用两个工具：
- Goreleaser：用于发布打包。您可以按照[此处](https://goreleaser.com/install/)的安装步骤操作；请注意匹配 [release.yml](/.github/workflows/release.yml) 中定义的版本
- golangci-lint：用于代码检查。您可以按照[此处](https://golangci-lint.run/usage/install/)的安装步骤操作；请注意匹配 [golangci-lint.yml](/.github/workflows/golangci-lint.yml) 中定义的版本

它们可以在每次推送或 PR 之前从仓库根目录执行：

**Goreleaser**
```shell
goreleaser build --snapshot --clean
```
**golangci-lint**
```shell
golangci-lint run
```

### 本地 NetBird 设置

> **重要**：以下所有步骤必须至少执行一次才能使开发环境启动并运行！

现在 NetBird 运行所需的一切都已安装，可以检出和设置实际的 NetBird 代码：

1. [Fork](https://guides.github.com/activities/forking/#fork) NetBird 仓库

2. 克隆您 fork 的仓库

   ```
   git clone https://github.com/<your_github_username>/netbird.git
   ```

3. 进入仓库文件夹

   ```
   cd netbird
   ```

4. 将原始 NetBird 仓库添加为 `upstream` 到您 fork 的仓库

   ```
   git remote add upstream https://github.com/netbirdio/netbird.git
   ```

5. 安装所有 Go 依赖项：

   ```
   go mod tidy
   ```

6. 配置 Git hooks 以进行自动代码检查：

   ```bash
   make setup-hooks
   ```

   这将配置 Git 在每次推送前自动运行代码检查，帮助及早发现问题。

### 开发容器支持

如果您更喜欢使用开发容器进行开发，NetBird 现在支持开发容器。
开发容器提供一致且隔离的开发环境，使贡献者更容易快速入门。请按照以下步骤在开发容器中设置 NetBird。

#### 1. 先决条件：

* 在您的机器上安装 Docker：[Docker 安装指南](https://docs.docker.com/get-docker/)
* 安装 Visual Studio Code：[VS Code 安装指南](https://code.visualstudio.com/download)
* 如果您更喜欢 JetBrains Goland，请遵循此[手册](https://www.jetbrains.com/help/go/connect-to-devcontainer.html)

#### 2. 克隆仓库：

按照前面的[本地 NetBird 设置](#本地-netbird-设置)克隆仓库。

#### 3. 在您选择的 IDE 中打开项目：

**VScode**：

在 Visual Studio Code 中打开项目文件夹：

```bash
code .
```

当您在 VS Code 中打开项目时，它会检测到开发容器配置的存在。
点击 VS Code 右下角的绿色"在容器中重新打开"按钮。

**Goland**：

打开 GoLand 并选择 `"文件" > "打开"` 以打开 NetBird 项目文件夹。
GoLand 将检测到开发容器配置并提示您在容器中打开项目。接受提示。

#### 4. 等待容器构建：

VSCode 或 GoLand 将使用指定的 Docker 镜像构建开发容器。这可能需要一些时间，具体取决于您的网络连接。

#### 6. 开发：

容器构建完成后，您可以在开发容器内开始开发。所有必需的依赖项和配置都在容器内设置完成。


### 构建和启动
#### 客户端

要启动 NetBird，请执行：
```
cd client
CGO_ENABLED=0 go build .
```

> Windows 客户端需要 Wireguard 驱动程序。您可以从 https://www.wintun.net/builds/wintun-0.14.1.zip 下载 wintun 驱动程序，解压后，您可以将文件 `windtun\bin\ARCH\wintun.dll` 复制到与您的二进制文件相同的路径或 `C:\Windows\System32\wintun.dll`。

> 要在 Windows 机器上使用 RDP 或虚拟化环境（例如 virtualbox 或云）测试客户端 GUI 应用程序，您需要从 https://fdossena.com/?p=mesa/index.frag 下载并提取 opengl32.dll 到构建的应用程序旁边。

在前台启动 NetBird 客户端：

```
sudo ./client up --log-level debug --log-file console
```
> 在 Windows 上使用具有管理员权限的 powershell
#### 信号服务

要启动 NetBird 的信号服务，请执行：

```
cd signal
go build .
```

要启动 NetBird 信号服务：

```
./signal run --log-level debug --log-file console
```

#### 管理服务
> 您可能需要为管理服务生成配置文件。请按照我们的[自托管指南](https://netbird.io/docs/getting-started/self-hosting)执行步骤 2 到 5。

要启动 NetBird 的管理服务，请执行：

```
cd management
go build .
```

要启动 NetBird 管理服务：

```
./management management --log-level debug --log-file console --config ./management.json
```

#### Windows Netbird 安装程序
创建 dist 目录
```shell
mkdir -p dist/netbird_windows_amd64
```

UI 客户端
```shell
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o netbird-ui.exe -ldflags "-s -w -H windowsgui" ./client/ui
mv netbird-ui.exe ./dist/netbird_windows_amd64/
```

客户端
```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o netbird.exe ./client/
mv netbird.exe ./dist/netbird_windows_amd64/
```
> Windows 客户端需要 Wireguard 驱动程序。您可以从 https://www.wintun.net/builds/wintun-0.14.1.zip 下载 wintun 驱动程序，解压后，您可以将文件 `windtun\bin\ARCH\wintun.dll` 复制到 `./dist/netbird_windows_amd64/`。

NSIS 编译器
- [Windows-nsis]( https://nsis.sourceforge.io/Download)
- [MacOS-makensis](https://formulae.brew.sh/formula/makensis#default)
- [Linux-makensis](https://manpages.ubuntu.com/manpages/trusty/man1/makensis.1.html)

NSIS 插件。下载并移动到 NSIS 插件文件夹。
- [EnVar](https://nsis.sourceforge.io/mediawiki/images/7/7f/EnVar_plugin.zip)
- [ShellExecAsUser](https://nsis.sourceforge.io/mediawiki/images/6/68/ShellExecAsUser_amd64-Unicode.7z)

Windows 安装程序
```shell
export APPVER=0.0.0.1
makensis -V4 client/installer.nsis
```

安装程序 `netbird-installer.exe` 将在根目录中创建。

### 测试套件

测试可以通过以下方式启动：

```
cd netbird
go test -exec sudo ./...
```
> 在 Windows 上使用具有管理员权限的 powershell

> 非 GTK 环境需要安装 `libayatana-appindicator3-dev`（debian/ubuntu）包

## 提交 PR 前的检查清单
作为关键网络服务和开源项目，我们在提交拉取请求之前必须强制执行一些事项：
- 保持函数尽可能简单，具有单一目的
- 尽可能使用私有函数和常量
- 为任何新的公共函数添加注释
- 为任何新的公共函数添加单元测试

> 在推送 PR 评论的修复时，请作为单独的提交推送；我们将在合并前压缩 PR，因此无需在推送前压缩，而且我们对单个 PR 中有 10-100 个提交完全没问题。这有助于审查对请求更改的修复。

## 其他项目仓库

NetBird 项目由 3 个主要仓库组成：
- NetBird：此仓库，包含代理和控制平面服务的代码
- Dashboard：https://github.com/netbirdio/dashboard，包含管理服务的管理 UI
- Documentations：https://github.com/netbirdio/docs，包含来自 https://netbird.io/docs 的文档

## 贡献者许可协议

为了避免以后出现任何潜在问题，很遗憾需要签署[贡献者许可协议](CONTRIBUTOR_LICENSE_AGREEMENT.md)。这可以简单地通过点击一个按钮来完成。

一旦 PR 被打开，机器人将自动发表评论要求签署协议。在签署之前，很遗憾无法合并。
