<div align="center">
<br/>
  <br/>
<p align="center">
  <img width="234" src="docs/media/logo-full.png"/>
</p>
  <p>
   <a href="https://img.shields.io/badge/license-BSD--3-blue)">
       <img src="https://sonarcloud.io/api/project_badges/measure?project=netbirdio_netbird&metric=alert_status" />
     </a> 
     <a href="https://github.com/netbirdio/netbird/blob/main/LICENSE">
       <img src="https://img.shields.io/badge/license-BSD--3-blue" />
     </a> 
    <br>
    <a href="https://docs.netbird.io/slack-url">
        <img src="https://img.shields.io/badge/slack-@netbird-red.svg?logo=slack"/>
     </a>
    <a href="https://forum.netbird.io">
        <img src="https://img.shields.io/badge/community forum-@netbird-red.svg?logo=discourse"/>
     </a>  
     <br>
    <a href="https://gurubase.io/g/netbird">
        <img src="https://img.shields.io/badge/Gurubase-Ask%20NetBird%20Guru-006BFF"/>
     </a>    
  </p>
</div>


<p align="center">
<strong>
  开始在 <a href="https://netbird.io/pricing">netbird.io</a> 使用 NetBird
  <br/>
  查看 <a href="https://netbird.io/docs/">文档</a>
  <br/>
   加入我们的 <a href="https://docs.netbird.io/slack-url">Slack 频道</a> 或 <a href="https://forum.netbird.io">社区论坛</a>
  <br/>
 
</strong>
<br>
<strong>
  🚀 <a href="https://careers.netbird.io">我们在招聘！加入 careers.netbird.io</a>
</strong>
<br>
<br>
<a href="https://registry.terraform.io/providers/netbirdio/netbird/latest">
    新增：NetBird Terraform 提供商
  </a> 
</p>

<br>

**NetBird 将无需配置的点到点专用网络与集中式访问控制系统结合在单一平台中，让您可以轻松为组织或家庭创建安全的专用网络。**

**连接。** NetBird 创建基于 WireGuard 的覆盖网络，自动通过加密隧道连接您的机器，无需繁琐的端口开放、复杂的防火墙规则、VPN 网关等操作。

**安全。** NetBird 通过精细的访问策略实现安全的远程访问，同时让您能够从一个直观的集中位置管理这些策略。适用于任何基础设施。

### 单一平台中的开源网络安全

https://github.com/user-attachments/assets/10cec749-bb56-4ab3-97af-4e38850108d2

### 自托管 NetBird（视频）
[![观看视频](https://img.youtube.com/vi/bZAgpT6nzaQ/0.jpg)](https://youtu.be/bZAgpT6nzaQ)

### 主要功能

| 连接性 | 管理 | 安全性 | 自动化 | 平台 |
|----|----|----|----|----|
| <ul><li>- \[x] 内核 WireGuard</ul></li> | <ul><li>- \[x] [管理 Web UI](https://github.com/netbirdio/dashboard)</ul></li> | <ul><li>- \[x] [SSO 和 MFA 支持](https://docs.netbird.io/how-to/installation#running-net-bird-with-sso-login)</ul></li> | <ul><li>- \[x] [公共 API](https://docs.netbird.io/api)</ul></li> | <ul><li>- \[x] Linux</ul></li> |
| <ul><li>- \[x] 点到点连接</ul></li> | <ul><li>- \[x] 自动对等发现和配置</ui></li> | <ul><li>- \[x] [访问控制 - 组和规则](https://docs.netbird.io/how-to/manage-network-access)</ui></li> | <ul><li>- \[x] [设置密钥用于批量网络部署](https://docs.netbird.io/how-to/register-machines-using-setup-keys)</ui></li> | <ul><li>- \[x] Mac</ui></li> |
| <ul><li>- \[x] 连接中继回退</ui></li> | <ul><li>- \[x] [IdP 集成](https://docs.netbird.io/selfhosted/identity-providers)</ui></li> | <ul><li>- \[x] [活动日志](https://docs.netbird.io/how-to/audit-events-logging)</ui></li> | <ul><li>- \[x] [自托管快速启动脚本](https://docs.netbird.io/selfhosted/selfhosted-quickstart)</ui></li> | <ul><li>- \[x] Windows</ui></li> |
| <ul><li>- \[x] [到外部网络的路由](https://docs.netbird.io/how-to/routing-traffic-to-private-networks)</ui></li> | <ul><li>- \[x] [私有 DNS](https://docs.netbird.io/how-to/manage-dns-in-your-network)</ui></li> | <ul><li>- \[x] [设备姿态检查](https://docs.netbird.io/how-to/manage-posture-checks)</ui></li> | <ul><li>- \[x] IdP 组与 JWT 同步</ui></li> | <ul><li>- \[x] Android</ui></li> |
| <ul><li>- \[x] 使用 BPF 的 NAT 穿透</ui></li> | <ul><li>- \[x] [多用户支持](https://docs.netbird.io/how-to/add-users-to-your-network)</ui></li> | <ul><li>- \[x] 点到点加密</ul></li> || <ul><li>- \[x] iOS</ui></li> |
||| <ul><li>- \[x] [使用 Rosenpass 的抗量子加密](https://netbird.io/knowledge-hub/the-first-quantum-resistant-mesh-vpn)</ui></li> || <ul><li>- \[x] OpenWRT</ui></li> |
||| <ul><li>- \[x] [定期重新认证](https://docs.netbird.io/how-to/enforce-periodic-user-authentication)</ui></li> || <ul><li>- \[x] [无服务器](https://docs.netbird.io/how-to/netbird-on-faas)</ui></li> |
||||| <ul><li>- \[x] Docker</ui></li> |

### NetBird 云端快速入门

- 在 [https://app.netbird.io/install](https://app.netbird.io/install) 下载并安装 NetBird
- 按照步骤使用 Google、Microsoft、GitHub 或您的电子邮件地址注册
- 查看 NetBird [管理 UI](https://app.netbird.io/)
- 添加更多机器

### 自托管 NetBird 快速入门

> 这是尝试自托管 NetBird 的最快方式。如果您已有公网域名和虚拟机，大约需要 5 分钟即可开始使用。
> 有关使用不同 IdP 的安装，请参阅[使用自定义身份提供商的高级指南](https://docs.netbird.io/selfhosted/selfhosted-guide#advanced-guide-with-a-custom-identity-provider)。

**基础设施要求：**
- 至少具有 **1 个 CPU** 和 **2GB** 内存的 Linux 虚拟机
- 该虚拟机应在 TCP 端口 **80** 和 **443** 以及 UDP 端口 **3478** 上可公开访问
- 指向该虚拟机的**公网域名**

**软件要求：**
- 虚拟机上安装 Docker 和 docker-compose 插件（[Docker 安装指南](https://docs.docker.com/engine/install/)），或使用 docker-compose 2.0 或更高版本
- 安装 [jq](https://jqlang.github.io/jq/)。在大多数发行版中通常可在官方仓库中找到，可以使用 `sudo apt install jq` 或 `sudo yum install jq` 安装
- 安装 [curl](https://curl.se/)。通常可在官方仓库中找到，可以使用 `sudo apt install curl` 或 `sudo yum install curl` 安装

**步骤**
- 下载并运行安装脚本：
```bash
export NETBIRD_DOMAIN=netbird.example.com; curl -fsSL https://github.com/netbirdio/netbird/releases/latest/download/getting-started.sh | bash
```
- 完成后，您可以通过 `docker-compose` 管理资源

### NetBird 内部原理简介
- 网络中的每台机器都运行 [NetBird 代理（或客户端）](client/)，用于管理 WireGuard
- 每个代理连接到[管理服务](management/)，该服务保存网络状态、管理对等 IP，并向代理（对等点）分发网络更新
- NetBird 代理使用 [pion/ice 库](https://github.com/pion/ice) 实现的 WebRTC ICE 来在建立机器之间的点到点连接时发现连接候选
- 连接候选通过 [STUN](https://en.wikipedia.org/wiki/STUN) 服务器的帮助发现
- 代理通过[信号服务](signal/)传递包含候选的 P2P 加密消息来协商连接
- 有时由于严格的 NAT（例如移动运营商级 NAT），NAT 穿透可能不成功，无法建立 P2P 连接。发生这种情况时，系统会回退到名为 [TURN](https://en.wikipedia.org/wiki/Traversal_Using_Relays_around_NAT) 的中继服务器，并通过 TURN 服务器建立安全的 WireGuard 隧道
 
在 NetBird 配置中成功用于 STUN 和 TURN 的是 [Coturn](https://github.com/coturn/coturn)。

<p float="left" align="middle">
  <img src="https://docs.netbird.io/docs-static/img/about-netbird/high-level-dia.png" width="700"/>
</p>

详细信息请参阅完整的[架构概述](https://docs.netbird.io/about-netbird/how-netbird-works#architecture)。

### 社区项目
-  [NetBird 安装脚本](https://github.com/physk/netbird-installer)
-  [Dominion Solutions 的 NetBird Ansible 集合](https://galaxy.ansible.com/ui/repo/published/dominion_solutions/netbird/)

**注意**：在开发过程中，`main` 分支可能处于不稳定甚至损坏的状态。
有关稳定版本，请参阅[发布](https://github.com/netbirdio/netbird/releases)。

### 支持致谢

2022 年 11 月，NetBird 加入了由德意志联邦共和国联邦教育与研究部赞助的 [StartUpSecure 计划](https://www.forschung-it-sicherheit-kommunikationssysteme.de/foerderung/bekanntmachungen/startup-secure)。NetBird 与 [CISPA 亥姆霍兹信息安全中心](https://cispa.de/en) 合作，将安全最佳实践和简单性带入专用网络。

![CISPA_Logo_BLACK_EN_RZ_RGB (1)](https://user-images.githubusercontent.com/700848/203091324-c6d311a0-22b5-4b05-a288-91cbc6cdcc46.png)

### 鸣谢
我们使用 [WireGuard®](https://www.wireguard.com/)、[Pion ICE (WebRTC)](https://github.com/pion/ice)、[Coturn](https://github.com/coturn/coturn) 和 [Rosenpass](https://rosenpass.eu) 等开源技术。我们非常欣赏这些开发者的工作，如果您能以任何方式支持他们（例如给个星标或贡献代码），我们将不胜感激。

### 法律声明
本仓库根据 BSD-3-Clause 许可证授权，适用于仓库的所有部分，但 management/、signal/ 和 relay/ 目录除外。
这些目录根据 GNU Affero 通用公共许可证 version 3.0 (AGPLv3) 授权。请参阅每个目录中的相应 LICENSE 文件。

_WireGuard_ 和 _WireGuard_ 徽标是 Jason A. Donenfeld 的[注册商标](https://www.wireguard.com/trademark-policy/)。
