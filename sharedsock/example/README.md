### 如何运行

这仅在 Linux 上有效

1. 运行 netcat 监听 UDP 端口 51820。这将是我们的外部进程：
```bash
nc -kluvw 1 51820
```

2. 构建并运行示例 Go 代码：

```bash
 go build -o sharedsock && sudo ./sharedsock
```

3. 通过发送 STUN 绑定请求测试逻辑

```bash
STUN_PACKET="000100002112A4425454" 
echo -n $STUN_PACKET | xxd -r -p | nc -u -w 1 localhost 51820
```

4. 您应该看到类似的 Go 程序输出。请注意，您也会在 netcat 服务器中看到一些二进制输出。这是因为内核将数据包复制到两个进程。

```bash 
 read a STUN packet of size 18 from ...
```

5. 发送非 STUN 数据包

```bash
echo -n 'hello' |  nc -u -w 1 localhost 51820
```

6. Go 程序不会打印任何内容。
