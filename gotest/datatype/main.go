package main

import "fmt"

var (
	a = 1
	b = a + 1
	c = b + a
)

func main() {
	fmt.Println(c)
	//数据类型
	var s = "hello 你好"
	var size = len(s)
	fmt.Println(size)
	for key, value := range s {
		fmt.Printf("key=%d,value=%#U\n", key, value)
	}
	var n = ' '
	fmt.Printf("空格=%d\n", n)
	var n1 = '你'
	fmt.Printf("你=%d\n", n1)

	const ss = "Go语言"
	for i, r := range ss {
		fmt.Printf("%#U  ： %d\n", r, i)
	}
}
