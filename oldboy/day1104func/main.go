package main

import "fmt"

/*
匿名函数：就是没有函数名的函数，定义格式如下
	func(参数)(返回值){
		函数体
	}
	由于没有函数名，所以没办法像普通函数那样调用，匿名函数需要保存到某个变量或者作为立即执行函数
	func main() {
		// 将匿名函数保存到变量
		add := func(x, y int) {
			fmt.Println(x + y)
		}
		add(10, 20) // 通过变量调用匿名函数

		//自执行函数：匿名函数定义完加()直接执行
		func(x, y int) {
			fmt.Println(x + y)
		}(10, 20)
	}

闭包：一个函数和与其相关的引用环境组合而成的实体
	闭包=函数+引用环境
*/
func main() {
	//将匿名函数保存到某个变量使用
	add := func(x int, y int) int {
		return x + y
	}
	sum := add(3, 4)
	fmt.Println("sum=", sum)

	//匿名函数立即执行
	result := func(x, y int) int {
		return x - y
	}(10, 4)
	fmt.Println("result=", result)

	//闭包使用
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f11 := adder()
	fmt.Println(f11(40)) //40
	fmt.Println(f11(50)) //90
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
