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

func main() {
	stu := Student{
		Age:  11,
		Name: "adsa",
	}
	fmt.Println(stu)
	fmt.Println(Student{Name: "dsds20", Age: 3})
	puiple := Puiple{stu}
	fmt.Println(puiple)
}
