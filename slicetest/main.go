package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(arr)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3, "len(s3)=", len(s3), "cap(s3)=", cap(s3))
	fmt.Println("s4=", s4, "len(s4)=", len(s4), "cap(s4)=", cap(s4))
	fmt.Println("s5=", s5, "len(s5)=", len(s5), "cap(s5)=", cap(s5))

	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	ss1 := []int{2, 4, 6, 8}
	fmt.Println(ss1)

	ss2 := make([]int, 16, 16)
	fmt.Println(ss2)

	copy(ss2, ss1)
	printSlice(ss2)
	ss2 = append(ss2[:3], ss2[4:]...)
	fmt.Println(ss2)
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}
