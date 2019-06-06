package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

var m = map[int]string{1: "a", 2: "b", 3: "c"}
var info string

type int8 int

func init() {
	var a int8 = 3333
	fmt.Println(a)
	fmt.Printf("map:%v\n", m)
	info = fmt.Sprintf("OS:%s,Arch:%s", runtime.GOOS, runtime.GOARCH)
}

func main() {
	fmt.Println(info)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入名字:")
	s, e := reader.ReadString('\n')
	if e != nil {
		fmt.Printf("fond an error:%s", e)
	} else {
		i := s[:len(s)-1]
		fmt.Printf("hello,%s!", i)
	}
	fmt.Println()
	var a [4]int
	var b = [4]int{}
	fmt.Printf("%v,%p,%T\n",a,&a,a)
	fmt.Printf("%v,%p,%T\n",b,&b,b)
}
