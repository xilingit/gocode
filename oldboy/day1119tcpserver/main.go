package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
	具体参考地址：https://www.liwenzhou.com/posts/Go/15_socket/
	OSI七层
		应用层---应用层---应用层、表示层、会话层
		传输层---传输层---传输层
		网络层---网络层---网络层
		网路接口层---数据链路层、物理层---数据链路层、物理层

	Go语言实现TCP通信
		TCP协议：是一种面向连接（连接导向）的、可靠的、基于字节流的传输层通信协议
		TCP服务端
			一个TCP服务端可以同时连接很多个客户端，Go语言创建多个goroutine实现并发非常方便高效
			每建立一次链接就创建一个goroutine去处理
			TCP服务端程序的处理流程:
				1、监听端口
				2、接收客服端请求建立链接
				3、创建goroutine处理链接

		出现粘包
			原因：TCP数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发
			解决办法：
				自己对数据包进行封包和拆包操作

*/
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10086") //监听端口
	if err != nil {
		fmt.Println("Listen failed,err:", err)
		return
	}
	for {
		conn, err := listener.Accept() //接收客服端请求建立链接
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn) //启动一个goroutine处理链接
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		fmt.Println("读到的n是", n)
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}
