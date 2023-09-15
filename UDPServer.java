import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;

public class UDPServer {
    private static void printUsage() {
        System.out.println("使用方法：java UDPServer port");
        System.out.println("例如：java UDPServer 30000");
    }
    public static void main(String[] args) {
        DatagramSocket socket = null;

        int port = 0;
        if(args.length != 1) {
            printUsage();
            return;
        } else {
            try {
                port = Integer.parseInt(args[0]);
            } catch (NumberFormatException e) {
                printUsage();
                return;
            }
            if(port > 65535) {
                printUsage();
                return;
            }
        }
        try {
            socket = new DatagramSocket(port);

            byte[] receiveData = new byte[1024];

            System.out.println("监听成功，按 Ctrl+c 结束");
            while (true) {
                // 创建数据包来接收客户端数据
                DatagramPacket receivePacket = new DatagramPacket(receiveData, receiveData.length);

                // 接收客户端数据
                socket.receive(receivePacket);

                // 从数据包中获取客户端发送的数据
                String clientMessage = new String(receivePacket.getData(), 0, receivePacket.getLength());
                System.out.println("接收到来自客户端的消息: " + clientMessage);

                // 获取客户端的地址和端口号，用于回复
                InetAddress clientAddress = receivePacket.getAddress();
                int clientPort = receivePacket.getPort();

                // 构建回复消息
                String responseMessage = "我收到了 " + clientMessage;
                byte[] responseData = responseMessage.getBytes();

                // 创建数据包来发送回复
                DatagramPacket responsePacket = new DatagramPacket(responseData, responseData.length, clientAddress, clientPort);

                // 发送回复
                socket.send(responsePacket);
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (socket != null && !socket.isClosed()) {
                socket.close();
            }
        }
    }
}
