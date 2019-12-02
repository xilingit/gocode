package main

import (
	"fmt"
	"os"
)

/*
	fmt包
	fmt包实现了类似C语言printf和scanf的格式化I/O，主要向外输出内容和输入内容两大部分

	Print系列函数会将内容输出到系统的标准输出
	Printf函数支持格式化输出字符串
	Println函数会在输出内容的结尾添加一个换行符

	Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，通常用这个函数往文件中写入内容

	Sprint系列函数会把传入的数据生成并返回一个字符串
*/
func main() {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错,err:", err)
		return
	}
	name := "哈哈哈"
	fmt.Fprintf(fileObj, "往文件里写入一下信息%s", name)
	fileObj.Close()
}
