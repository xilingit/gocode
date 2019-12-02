package main

import "fmt"

/*
	接口定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节

	接口(interface)是一种抽象类型，是一组method的集合
	定义格式如下:
		type 接口类型名 interface{
			方法名1(参数列表1) 返回值列表1
			...
		}
		接口名，一般会在单词后面加上er。如Writer、Stringer

	实现接口的条件
		一个对象只要全部实现了接口中的方法，那么就实现了这个接口

	接口类型变量
		能够存储所有实现了该接口的实例


	值接收者实现接口：不管结构体还是结构体指针类型都能赋值给接口变量
	指针接收者实现接口：只能接收结构体指针类型变量

	一个类型可以实现多个接口

	接口嵌套
		接口与接口间可以通过嵌套创造出新的接口
			type Sayer interface{
				say()
			}
			type Mover interface{
				move()
			}
			type animal interface{
				Sayer
				Mover
			}

	空接口
		空接口是没有定义任何方法的接口，因此任何类型都实现了空接口
		空接口类型的变量可以存储任意类型的变量
		var x interface{}
		x = ...
		作用：
			空接口作为函数的参数
			空接口作为map的值

	类型断言
		x.(T)
		x:表示类型为interface{}的变量
		T:表示断言x可能是的类型
		返回两个参数，第一个参数是x转化为T类型后的变量，第二个是一个布尔值，
		若为true则表示断言成功，为false则表示断言失败

*/

type Sayer interface {
	say() //无参无返回值
}

type Cat struct{}

func (c Cat) say() { //实现接口
	fmt.Println("喵喵喵")
}

type Dog struct{}

func (d *Dog) say() { //实现接口
	fmt.Println("汪汪汪")
}

func main() {
	var say Sayer

	//以下二者都可以
	say = Cat{}
	say = &Cat{} //Go语言中有对指针类型变量求值的语法糖,内部会自动求值*Cat()
	say.say()

	//注意这个
	// say = Dog{}//会报错
	say = &Dog{}
	say.say()

	//空接口
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x) //string
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x) //int
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x) //bool
	var m = make(map[string]interface{})
	m["name"] = "hello"
	m["age"] = 18
	m["married"] = false
	fmt.Println(m)

	//类型断言
	v, ok := x.(bool)
	if ok {
		fmt.Printf("它是个布尔类型，值为%v\n", v)
	} else {
		fmt.Println("断言失败")
	}
	justifyType(x)
}

func justifyType(i interface{}) {
	switch v := i.(type) { //多用switch来做类型断言
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
