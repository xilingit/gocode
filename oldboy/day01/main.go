package main

import (
	"fmt"
)

//声明变量
var name string
var age int
var isOk bool
var (
	a = 1
	b = 3
	c = "dds"
)

func main() {
	fmt.Println("hello world")
	name = "tom"
	age = 22
	isOk = true
	fmt.Println(name, age, isOk)
	fmt.Println(a, b, c)
	//声明变量同时赋值
	var s1 string = "王"
	//类型推导，无法修改成其他类型的值
	var s2 = "hai"
	//简短变量声明，只能用在函数中，无法用在全局变量中
	s3 := true
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
