## 模拟对象

要从 iface 包接口生成（或刷新）模拟对象，请安装 [mockgen](https://github.com/golang/mock)。
运行此命令以更新 PacketFilter 模拟对象：
```bash
mockgen -destination iface/mocks/filter.go -package mocks github.com/netbirdio/netbird/iface PacketFilter
```
