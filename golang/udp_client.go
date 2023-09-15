package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 检查启动参数是否正确
	if len(os.Args) != 4 {
		fmt.Println("Usage: udp_client <server_ip> <server_port> <message>")
		os.Exit(1)
	}

	serverIP := os.Args[1]
	serverPort := os.Args[2]
	message := os.Args[3]

	// 解析服务器地址
	serverAddr := serverIP + ":" + serverPort
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		os.Exit(1)
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 发送消息到服务器
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Printf("Sent message to server: %s\n", message)

	// 接收服务器的响应
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		os.Exit(1)
	}

	serverResponse := string(buffer[:n])
	fmt.Println("Server response:", serverResponse)
}
