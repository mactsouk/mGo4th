package main

import (
	"fmt"
)

func InitSliceNew(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func InitSliceAppend(n int) []int {
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func main() {
	fmt.Println(InitSliceNew(10))
	fmt.Println(InitSliceAppend(10))
}
