package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
	TCP客服端
	一个客服端进行TCP通信的流程如下
		1、建立与服务端的链接
		2、进行数据收发
		3、关闭链接
*/
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("connect err,err msg is", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入,\n是界定符
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
