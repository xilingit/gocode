package main

import (
	"fmt"
)

func main() {
	nums := [3]int{}
	nums[0] = 1
	dnums := nums[0:2]
	dnums = append(dnums, []int{2, 3}...)
	nums[0] = 5
	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))

	dst := []int{1, 2, 3}
	src := []int{4, 5, 6, 7, 8}
	n := copy(dst, src)

	fmt.Printf("dst: %v, n: %d", dst, n)
}
