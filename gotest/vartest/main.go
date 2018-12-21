package main

import "fmt"

var x, y, z int8

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
const q, w, r = 2, true, 2.2

var p, o, i = 1, true, 22.0

func main() {
	//g, h := 123, "hello"
	//println(x, y, a, b, c, d, e, f, g, h)
	x = 4
	a := &x
	fmt.Printf("a的类型是%T\n", a)
	fmt.Printf("a中存的值是%d\n", *a)
	fmt.Printf("x地址是%p y地址是%p z地址是%p\n", &x, &y, &z)
}
