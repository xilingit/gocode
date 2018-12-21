package main

import "fmt"

func main() {
	var mySlice []int = make([]int, 4, 10)
	//mySlice = [] int {3,2}
	mySlice[0] = 1
	fmt.Println(mySlice)
	str := "hello,世界"
	fmt.Printf("%T\n", str[2])
	mySlice = append(mySlice, 4)
	fmt.Println(mySlice)

	var map1 = map[int]string{
		2: "hello",
		3: "world",
		1: "😄",
	}
	fmt.Println(map1)
	delete(map1, 4)
	fmt.Println(map1)

	f := func(a, b int, z float64) bool {
		return a*b < int(z)
	}

	fmt.Printf("f的类型是%T", f)
}
