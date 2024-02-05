package main

import (
	"fmt"
)

func main() {
	for x := range 5 {
		fmt.Println(x)
	}

	// go 1.22
	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("%d ", val)
		}()
	}
}
