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
	str := make([]byte, 3)

	nChar, err := Read(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read %d random bytes\n", nChar)
		fmt.Printf("The 3 random bytes are: %v\n", str)
	}

	// Getting a random int value
	var max int = 100
	n := rand.N(max)
	fmt.Println("integer n =", n)

	// Getting a random uint value
	var uMax uint = 100
	uN := rand.N(uMax)
	fmt.Println("unsigned int uN =", uN)
}
