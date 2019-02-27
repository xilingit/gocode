package main

import (
	"fmt"
	"unicode/utf8"
)

type (
	City string
	B0   int8
	B1   int16
	B2   int32
	B3   int64
)

func main() {
	city := City("sda")
	fmt.Println(city)
	fmt.Println(B0(2))
	fmt.Println(B1(2))
	fmt.Println(B3(2))
	fmt.Println(B2(2) / B2(B0(2)))
	n := utf8.RuneCountInString("😂😄")
	fmt.Println(n)
}
