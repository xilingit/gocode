package main

import "fmt"

/*
go语言中的引用类型
	指针、切片、map、channel
*/
func main() {
	/*
		go语言中的指针不能进行偏移和运算，是安全指针
		go语言中的函数传参都是值拷贝
		&取地址		*根据地址取值
		go语言中的值类型（int、float、bool、string、array、struct）都有对应
		的指针类型
	*/
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	/*
		new和make都是内建函数
		在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间
		否则我们的值就没法存储
		对于值类型的声明不需要分配内存空间，是因为它们在声明的时候就已经默认分配好了内存空间
	*/

	/*
		new函数  	func new(Type) *Type
	*/
	var e *int //只是声明，但是没有初始化
	e = new(int)
	d := new(bool)
	*e = 10
	*d = true
	fmt.Println(e)
	fmt.Println(d)
	fmt.Println(*e)
	fmt.Println(*d)

	/*
		make函数  	func make(t Type,size ... IntegerType) Type
		它只用于slice、map以及chan的内存创建，返回类型就是这三个类型本身
	*/
	var f map[string]int
	f = make(map[string]int, 10)
	f["关"] = 33
	fmt.Println(f)
	fmt.Println(f["关"])
	/*
		new与make的区别
		1、二者都是用来做内存分配的
		2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
		3、而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	*/
}
