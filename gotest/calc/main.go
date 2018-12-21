package main

import "fmt"

func main() {
	var a int = 4
	var ptr *int

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Println(ptr)
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
