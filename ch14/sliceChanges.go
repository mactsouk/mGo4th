package main

import (
	"fmt"
	"slices"
)

func main() {

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := []int{5, 6}
	res := slices.Concat(s1, s2, s3)
	fmt.Println(res)

	// go 1.22
	src := []int{11, 12, 13, 14}
	// delete #1 and #2
	mod := slices.Delete(src, 1, 3)
	fmt.Println("src:", src)
	fmt.Println("mod:", mod)
}
