package main

import (
	"fmt"
	"time"
)

func createSlice() []int {
	return make([]int, 1000000)
}

func getValue(s []int) []int {
	returnVal := make([]int, 3)
	copy(returnVal, s)
	return returnVal
}

func main() {
	for i := 0; i < 15; i++ {
		message := createSlice()
		val := getValue(message)
		fmt.Print(len(val), " ")
		time.Sleep(10 * time.Millisecond)
	}
}
