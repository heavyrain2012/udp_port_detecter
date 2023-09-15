## 环境要求
需要java环境

## 服务器端
把项目中的[UDPServer.class](./UDPServer.class)拷贝到服务器上，然后执行：
```
java UDPServer 30000
```
这样就启动了端口30000的监听。

## 客户端
把[UDPClient.class](./UDPClient.class)拷贝到电脑上，在命令行执行
```
java UDPClient 192.168.1.120 30000 hello
```
这样客户端就会连指定ip的3000端口并发送hello，如果显示：“我收到了 hello”，说明端口是开放的。实际使用中把IP地址换为服务器的公网IP地址。
