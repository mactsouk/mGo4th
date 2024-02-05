package main

import (
	"fmt"
	"math/rand/v2"
)

func Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); {
		val := rand.Uint64()
		for j := 0; j < 8 && i < len(p); j++ {
			p[i] = byte(val & 0xff)
			val >>= 8
			i++
		}
	}
	return len(p), nil
}

func main() {
	b := make([]byte, 5)
	Read(b)
	fmt.Printf("5 random bytes: %v\n", b)

	{
		// random integer
		var max int = 100
		n := rand.N(max)
		fmt.Println("integer n =", n)
	}

	{
		// random unsigned integer
		var max uint = 100
		n := rand.N(max)
		fmt.Println("unsigned int n =", n)
	}
}
