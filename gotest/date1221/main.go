package main

import (
	"fmt"
)

type Student struct {
	Age  int
	Name string
}
type Puiple struct {
	Student //匿名成员
}
type Tree struct {
	value       int
	left, right *Tree
}

func main() {
	stu := Student{
		Age:  11,
		Name: "adsa",
	}
	fmt.Println(stu)
	fmt.Println(Student{Name: "dsds20", Age: 3})
	puiple := Puiple{stu}
	fmt.Println(puiple)
	tree1 := Tree{
		value: 1,
	}
	tree2 := Tree{
		value: 2,
	}
	fmt.Printf(">>>%#v\n", tree1 == tree2)
	i := sum(2, 4, 1, 3, 4, 5, 8, 9)
	fmt.Println("i=", i)
	func(name string) {
		fmt.Println(name)
	}("hebei")

	//闭包
	addOne := addInt(1)
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())

	addTwo := addInt(2)
	fmt.Println(addTwo())
	fmt.Println(addTwo())
	fmt.Println(addTwo())
}

func addInt(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

func plus(a, b int) int {
	return a + b
}

func jiaJian(aa, bb int) (sum, des int) {
	sum = aa + bb
	des = aa - bb
	return
}

func sum(nums ...int) int {
	fmt.Println("len of nums is :", len(nums))
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
