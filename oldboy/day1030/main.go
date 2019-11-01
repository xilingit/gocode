package main

import (
	"fmt"
)

func main() {
	// math.MaxFloat32 //float32最大值
	f1 := 1.23456 //默认Go语言中的小数都是float64类型
	fmt.Printf("%T\n", f1)
	f2 := float32(1.246464) //显示声明float32类型的变量
	fmt.Printf("%T\n", f2)

	s := "hello沙河"
	n := len(s)
	fmt.Println((n))
	for i, v := range s { //从字符串中拿出具体的字符
		fmt.Printf("i=%v,v=%c\n", i, v)
		fmt.Println()
	}
	v := 0b01010
	fmt.Println(v)
	changeString()

	finger := 3
	switch finger {
	default:
		fmt.Println("无效的输入！")
	case 0 + 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	}
	s1 := "a"
	switch {
	case s1 == "a":
		fmt.Println("a")
		fallthrough
	case s1 == "b":
		fmt.Println("b")
	case s1 == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
