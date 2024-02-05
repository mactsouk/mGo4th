package main

import (
	"fmt"
	"time"
)

func main() {
	for x := range 5 {
		fmt.Print(" ", x)
	}
	fmt.Println()

	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("%d ", val)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}

