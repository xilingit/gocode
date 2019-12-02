package main

import "fmt"

/*
	defer语句
	Go语言中的defer语句会将其后面跟随的语句进行延迟处理
	先被defer的语句最后被执行，最后被defer的语句，最先被执行
	一般用于资源清理、文件关闭、解锁及记录时间等

	参考:https://www.liwenzhou.com/posts/Go/09_function/，defer执行时机

*/
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/*  defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
输出结果是
	A 1 2 3
	B 10 2 12
	BB 10 12 22
	AA 1 3 4
*/

/*
先确定14行defer calc("AA", x, calc("A", x, y))中calc("A", x, y)的值
	A 1 2 3
再确定16行defer calc("BB", x, calc("B", x, y))中calc("B", x, y)的值
	B 10 2 12
再执行defer calc("BB", 10, 12)
	BB 10 12 22
最后执行defer calc("AA", 1, 3)
	AA 1 3 4
*/
