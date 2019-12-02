package main

import (
	"fmt"
	"strconv"
)

/*
	Go语言内置包之strconv
	实现了基本数据类型和其字符串表示的相互转换，主要有以下常用函数:
	Atoi()、Itia()、parse系列、format系列、append系列

	string与int类型转换
		将字符串类型的整数转换为int类型
		func Atoi(s string) (i int, err error)

		将int类型数据转换为对应的字符串表示
		func Itoa(i int) string

	Parse系列函数
		Parse类函数用于转换字符串为给定类型的值

		func ParseBool(str string) (value bool, err error)
		它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

		func ParseInt(s string, base int, bitSize int) (i int64, err error)

		func ParseUint(s string, base int, bitSize int) (n uint64, err error)

		func ParseFloat(s string, bitSize int) (f float64, err error)

	Format系列函数
		实现了将给定类型数据格式化为string类型数据的功能
*/
func main() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-2", 10, 64)
	u, _ := strconv.ParseUint("10", 10, 64)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(u)
}
