package main

import "fmt"

func main() {
	const (
		Sunday = iota
		_
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		nuberOfDays
	)
	fmt.Println(nuberOfDays)
	fmt.Println(Tuesday)
	var str string = "hello,中国"
	ch := str[1]
	fmt.Printf("ch的类型是%T,ch=%c\n", ch, ch)
	for k, v := range str {
		fmt.Printf("k=%d,v=%c\n", k, v)
	}
	var b byte = 0
	fmt.Println(b)
	const N = 3
	var c [2 * N]int
	c[0] = 3
	c[1] = 31
	c[2] = 32
	c[3] = 4
	c[4] = 40
	c[5] = 42
	fmt.Println(c)
}
