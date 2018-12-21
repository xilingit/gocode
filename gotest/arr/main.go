package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 2
	a[1] = 0
	a[2] = 6
	a[3] = 5
	var b = []float32{2, 3.3, 4.1}
	fmt.Println(a)
	float32s := append(b, 333)
	fmt.Println(float32s)
	fmt.Printf("b[0] type is %T\n", b[0])
	fmt.Println(a[0])

	var a1 = 20  /* 声明实际变量 */
	var ip1 *int /* 声明指针变量 */

	ip1 = &a1 /* 指针变量的存储地址 */

	fmt.Printf("a1 变量的地址是: %x\n", &a1)

	/* 指针变量的存储地址 */
	fmt.Printf("ip1 变量储存的指针地址: %x\n", ip1)

	/* 使用指针访问值 */
	fmt.Printf("*ip1 变量的值: %d\n", *ip1)

	var p *int
	fmt.Printf("ptr 的值为 : %x\n", p)
}
