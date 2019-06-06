package main

import (
	"fmt"
	"strconv"
	"strings"
)

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a, b int) int {
	return op(a, b)
}

func main() {
	for i := 100; i <= 999; i++ {
		if isNumber1(strconv.Itoa(i)) {
			fmt.Println(i, "is 水仙花")
		}
	}
	sum := f(4)
	fmt.Println(sum)

	s := urlProcess("www.baidu.com")
	fmt.Println(s)

	c := add
	fmt.Println(c)
	result := operator(c, 100, 2)
	fmt.Println(result)
	deferTest()
	wanshu(1000)
}

func wanshu(n int) {
	for i := 0; i < n; i++ {
		if isWanshu(i) {
			fmt.Println(i)
		}
	}
}

func isWanshu(n int) bool {
	var sum int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return n == sum
}

func deferTest() {
	var i = 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

func urlProcess(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

//求阶乘
func f(n int) uint64 {
	var sum1 uint64
	for j := 1; j <= n; j++ {
		var sum uint64 = 1
		for i := 1; i <= j; i++ {
			sum *= uint64(i)
		}
		sum1 += sum
	}
	return sum1
}

//判断是否是水仙花
func isNumber(n int) bool {
	var i, j, k int //i 百位  j 十位  k 个位
	k = n % 10      //个位
	j = (n / 10) % 10
	i = n / 100
	//fmt.Println(i, j, k)
	return n == i*i*i+j*j*j+k*k*k
}

//判断是否是水仙花
func isNumber1(s string) bool {
	var sum int
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')
		sum += num * num * num
	}
	number, err := strconv.Atoi(s)
	if err != nil {
		//fmt.Println("can not convert %s to int",s)
		return false
	}
	return sum == number
}
