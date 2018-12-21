package main

import "fmt"

func main() {
	sum, sub := sumAndSub(33, 55)
	fmt.Println(sum, sub)
}
func sumAndSub(a int, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
