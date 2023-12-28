package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func initializeValue() {
	x = 5
}

func main() {
	function := sync.OnceFunc(initializeValue)

	for i := 0; i < 10; i++ {
		go function()
	}
	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		x = x + 1
	}
	fmt.Printf("x = %d\n", x)

	for i := 0; i < 10; i++ {
		go function()
	}
	time.Sleep(time.Second)
	fmt.Printf("x = %d\n", x)
}
