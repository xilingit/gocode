package main

import "fmt"

func main() {
	//var a = 1
	//var b = int64(1)
	//fmt.Printf("a=%v,a type is %T\n", a, a)
	//fmt.Printf("b=%v,b type is %T\n", b, b)
	//
	//const (
	//	c = 1
	//	d,e = true,"asd"
	//)
	//
	//var i int64 = 22
	//var j int8
	//j = int8(i)+127
	//fmt.Println(j)
	//var s = `dsdadsd`
	//bb := s[0]
	//fmt.Printf("bb type is %T,bb = %c\n",bb,bb)
	//
	//var str = "hello, 世界"
	//
	//for _, char := range str {
	//	fmt.Printf("%T,char=%c\n", char,char)
	//}

	//arr := [5]int{1, 2, 3, 4, 5}
	//slice := arr[0:3] // 左闭右开区间，最终切片为 [1,2,3]
	//fmt.Printf("slice=%v,len(slice)=%d,cap(slice)=%d\n", slice, len(slice), cap(slice))
	//slice2 := make([] int, 6)
	//slice = append(slice, slice2...)
	//fmt.Printf("len(slice)=%d,cap(slice)=%d\n", len(slice), cap(slice))

	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[0:3]
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	newSlice[0] = 6
	fmt.Println("after modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
}
