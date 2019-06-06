package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println(format)
}
