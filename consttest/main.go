package main

import "fmt"

const (
	Unknown = iota
	Female  = "ha"
	Male
	WECHAT = iota + 1
	Alipay
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println(Unknown) //0
	fmt.Println(Female)  //ha
	fmt.Println(Male)    //ha 保留上一行数据
	fmt.Println(WECHAT)  //4
	fmt.Println(Alipay)  //5 保留上一行数据
	fmt.Println("i=", i) //1
	fmt.Println("j=", j) //6
	fmt.Println("k=", k) //12
	fmt.Println("l=", l) //24
}
