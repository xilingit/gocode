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

//批量声明常量
const (
	OK       = 200
	notFound = 404
)

const (
	n1 = 100
	n2 //默认和上边一致
	n3 //默认和上边一致
)

//iota 计数器
const (
	a1 = iota //0
	a2 = iota //1
	a3        //2
)

const (
	b1 = iota //0
	b2 = iota //1
	_  = iota //2
	b3 = iota //3
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1:1  d2:2
	d3, d4 = iota + 1, iota + 2 //d3:2  d4:3
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
	fmt.Println("a1=", a1)
	fmt.Println("a2=", a2)
	fmt.Println("a3=", a3)
	fmt.Println("KB=", KB)

	fmt.Println()

	//10进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("i1的%%b=%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%X\n", i1)
	fmt.Printf("%T\n", i1)
	var i2 int8 = 9
	i2 = int8(22)
	fmt.Println(i2)
}
