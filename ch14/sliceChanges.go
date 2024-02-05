package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2}
	s2 := []int{-1, -2}
	s3 := []int{10, 20}
	conCat := slices.Concat(s1, s2, s3)
	fmt.Println(conCat)

	v1 := []int{-1, 1, 2, 3, 4}
	fmt.Println("v1:", v1)
	v2 := slices.Delete(v1, 1, 3)
	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)
}
