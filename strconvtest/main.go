package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := strconv.IntSize
	fmt.Println(i)
	e := strconv.ErrRange
	fmt.Println(e)
	b, e := strconv.ParseBool("1")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(b)
	}
	var s string = "hello"
	s = s[1:]
	fmt.Println(s)
	u, e := strconv.ParseUint("11", 3, 0)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(u)
	}
	u1, e := strconv.ParseInt("-11", 2, 0)
	if e != nil {
		fmt.Println(e)
		fmt.Println(u1)
	} else {
		fmt.Println(u1)
	}
	fmt.Println(strconv.ParseBool("1")) // true
	var a = 'E'
	fmt.Printf("%d", a)

	sr, err := strconv.Unquote(`"\"大\t家\t好！\""`)
	fmt.Println(sr, err)
	sr, err = strconv.Unquote(`'大家好！'`)
	fmt.Println(sr, err)
	sr, err = strconv.Unquote(`'好'`)
	fmt.Println(sr, err)
	sr, err = strconv.Unquote("`大\\t家\\t好！`")
	fmt.Println(sr, err)
}
