package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", "192.168.20.23:8888")
	if e != nil {
		log.Fatal(e)
	}
	defer listener.Close()
	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Fatal(e)
		}
		fmt.Printf("访问客户端信息： con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()//关闭conn

	for {

		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", c.RemoteAddr().String())
		buf := make([]byte, 1024 )
		n, err := c.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}

		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}
