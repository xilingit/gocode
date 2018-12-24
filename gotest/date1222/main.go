package main

import "fmt"

func main() {
	//97天，几个星期几天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("97天是%d个星期%d天\n", week, day)

	//定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：
	// 5/9*（华氏温度-100），请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度%.2f对应的摄氏温度是%.2f\n", huashi, sheshi)

	//--------------------------------
	//演示逻辑运算符的使用
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}
}
