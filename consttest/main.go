package main

import "fmt"

const (
	Unknown = iota
	Female  = "hahhha"
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
	fmt.Println(Unknown)
	fmt.Println(Female)
	fmt.Println(Male)
	fmt.Println(WECHAT)
	fmt.Println(Alipay)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
