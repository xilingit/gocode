package main

import "fmt"

/*
go语言中支持函数、匿名函数和闭包
func 函数名(参数)(返回值){
	函数体
}

我们可以使用type关键字来定义一个函数类型，具体格式如下:
	type calculation func(int,int) int
上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值
简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
	func add(x, y int) int {
		return x + y
	}

	func sub(x, y int) int {
		return x - y
	}
add和sub都能赋值给calculation类型的变量。
	var c calculation
	c = add

函数可以作为参数
	func add(x, y int) int {
		return x + y
	}
	func calc(x, y int, op func(int, int) int) int {
		return op(x, y)
	}
	func main() {
		ret2 := calc(10, 20, add)//add函数作为参数
		fmt.Println(ret2) //30
	}
函数作为返回值
	func do(s string) (func(int, int) int, error) {
		switch s {
		case "+":
			return add, nil
		case "-":
			return sub, nil
		default:
			err := errors.New("无法识别的操作符")
			return nil, err
		}
	}
*/
type cal func(int, int) int

func main() {
	result := sum(3, 4)
	fmt.Println(result)
	result2 := sum2(3, 4, 5, 6, 7)
	fmt.Println(result2)
	var c cal
	c = sum
	cSum := c(2, 5)
	fmt.Println(cSum)
	e := decrise
	cDes := e(5, 2)
	fmt.Println(cDes)
}

//求两个数之和
func sum(x, y int) int {
	return x + y
}

//求两个数的差
func decrise(x, y int) int {
	return x - y
}

func sum2(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
