package main

import (
	"flag"
	"fmt"
)

/*
	Go语言内置的flag包实现了命令行参数的解析，可以用来开发命令行工具
		os.Args是一个[]string
		go run main.go a b c d

	定义命令行flag参数
		有以下两种常用的定义命令行flag参数的方法
		1、flag.Type()
				name := flag.String("name", "张三", "姓名")
				此时name为指针类型
		2、flag.TypeVar()
				var name string
				flag.StringVar(&name, "name", "张三", "姓名")
		通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析
*/

func main() {
	var age int
	name := flag.String("name", "san", "姓名")
	flag.IntVar(&age, "age", 10, "年龄")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(age)
}
