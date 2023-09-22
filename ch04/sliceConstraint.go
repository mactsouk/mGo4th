package main

import (
	"fmt"
)

func f1[S interface{ ~[]E }, E interface{}](x S) int {
	return len(x)
}

func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}

func f3[S ~[]E, E any](x S) int {
	return len(x)
}

func main() {
	fmt.Println("Len:", f1([]int{1, 2, 3}))
	fmt.Println("Len:", f2([]float64{1.1, -2}))
	fmt.Println("Len:", f3([]float32{1.1, -2}))
}
