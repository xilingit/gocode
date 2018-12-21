package main

import (
	"fmt"
)

func main() {
	//fmt.Println(len("a"))
	//fmt.Println(len("你好"))
	//s := fmt.Sprintf("%d:%s", 2018, "年")
	//fmt.Println(s)
	//
	//var b1 strings.Builder
	//b1.WriteString("aaa")
	//b1.WriteString(" and ")
	//b1.WriteString("ccc")
	//fmt.Println(b1.String())

	a := 22
	b := new(int)
	fmt.Printf("b的类型是%T\n", b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("*b =", *b)

	var c uint = +0
	fmt.Println(c)
}
