# UDP端口探测
野火音视频常见客户问题是UDP端口的联通性问题，这里使用UDP端口探测工具来验证野火音视频服务和TURN服务是否开放了UDP端口。

## 工作原理
有2个工具，服务器端和客户端，服务器端监听指定的UDP端口，客户端去给这个端口发送消息，服务器端收到消息并回复。当客户端能够收到响应，说明此服务器端口是通的，然后再尝试下一个端口，以此类推。

## 服务器端
服务器端有2个版本，一个是golang版本，另外一个是java版本。golang版本支持window+amd64/linux+amd64/linux_arm64/osx+amd64/osx+arm64平台，如果是这些平台的服务器，建议用golang版本。其他平台请使用java版本。具体使用方法，请参考下面目录：
1. [golang版本](./golang)
2. [java版本](./java)

## 客户端
客户端有3个版本，golang版本、java版本和android版本。golang和java版本是运行在电脑上的，android版本是运行在android手机上的。golang版本支持window+amd64/linux+amd64/linux_arm64/osx+amd64/osx+arm64平台，如果是这些平台的服务器，建议用golang版本。其他电脑平台请使用java版本。具体使用方法请参考下面目录：
1. [golang版本](./golang)
2. [java版本](./java)
2. [android版本](./android)
