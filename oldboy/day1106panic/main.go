package main

import "fmt"

/*
	Go语言中目前没有异常机制，使用panic/recover来处理
	panic可以在任何地方引发，但recover只有在defer调用的函数中有效

	注意
		1、recover()必须搭配defer使用
		2、defer一定要在可能引发panic的语句之前定义
*/
func main() {
	funcA()
	funcB()
	funcC()
}
func funcA() {
	fmt.Println("funcA")
}
func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发生异常了")
		}
	}()
	panic("funcB")
}
func funcC() {
	fmt.Println("funcC")
}
