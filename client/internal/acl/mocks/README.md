## 模拟对象

要从 acl 包生成（或刷新）模拟对象，请安装 [mockgen](https://github.com/golang/mock)。
从 `./client/internal/acl` 文件夹运行此命令以更新 iface mapper 接口模拟对象：
```bash
mockgen -destination mocks/iface_mapper.go -package mocks . IFaceMapper
```
