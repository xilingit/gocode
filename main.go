package main

import "fmt"

func Test(a *[3]int8){
	a[1] = 33
}

func main(){
	arr := [3]int8{2,3,4}
	fmt.Println(arr)
	Test(&arr)
	fmt.Println(arr)
}
