package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 检查启动参数是否包含端口号
	if len(os.Args) != 2 {
		fmt.Println("Usage: udp_server <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// 创建UDP连接
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on port", port)

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}

		clientMessage := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", addr.IP, clientMessage)

		responseMessage := "我收到了 " + clientMessage
		_, err = conn.WriteToUDP([]byte(responseMessage), addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}
