package main

import "fmt"

func main() {
	//切片
	//数组的长度是固定并且数组长度属于类型的一部分，所以数组有很多的局限性
	//切片是一个拥有相同类型元素的可变长度的序列，它是基于数组类型做的一层封装
	//它非常灵活，支持自动扩容
	//切片是一个引用类型，它的内部结构包含地址、长度和容量
	// var a []int        //切片声明
	// a = []int{2, 3, 4} //赋值
	// fmt.Println(a)
	//切片拥有自己的长度和容量，使用内置的len()函数求长度，使用cap()函数求切片的容量

	//可以基于数组定义切片
	b := [5]int{2, 4, 5, 6, 1}
	c := b[1:5]
	fmt.Println(c)

	var a = make([]string, 5, 10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	for i, v := range a {
		fmt.Printf("%d=%#v\n", i, v)
	}
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
