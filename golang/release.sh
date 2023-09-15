GOOS=windows GOARCH=amd64 go build -o release/server/windows/amd64/UdpServer.exe udp_server.go
GOOS=linux GOARCH=amd64 go build -o release/server/linux/amd64/UdpServer udp_server.go
GOOS=linux GOARCH=arm64 go build -o release/server/linux/arm64/UdpServer udp_server.go
GOOS=darwin GOARCH=amd64 go build -o release/server/darwin/amd64/UdpServer udp_server.go
GOOS=darwin GOARCH=arm64 go build -o release/server/darwin/arm64/UdpServer udp_server.go


GOOS=windows GOARCH=amd64 go build -o release/client/windows/amd64/UdpClient.exe udp_client.go
GOOS=linux GOARCH=amd64 go build -o release/client/linux/amd64/UdpClient udp_client.go
GOOS=linux GOARCH=arm64 go build -o release/client/linux/arm64/UdpClient udp_client.go
GOOS=darwin GOARCH=amd64 go build -o release/client/darwin/amd64/UdpClient udp_client.go
GOOS=darwin GOARCH=arm64 go build -o release/client/darwin/arm64/UdpClient udp_client.go