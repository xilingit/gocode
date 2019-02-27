package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, e := ioutil.ReadFile("./mytestgo/gotest/filetest/abc.txt")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(string(file))
	}
}
