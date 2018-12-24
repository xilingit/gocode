package main

import "fmt"

func main() {
	//a b两个变量，不实用临时变量进行交换
	var a = 20
	var b = 30
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	c := 100
	var ptr = &c
	fmt.Println(ptr)

}
