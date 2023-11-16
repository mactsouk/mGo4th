package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	times := 10
	seed := int64(0)
	arguments := os.Args
	if len(arguments) > 1 {
		n, err := strconv.ParseInt(arguments[1], 10, 64)
		if err == nil {
			fmt.Println(n)
			seed = n
		}
	}

	fmt.Println("Using seed:", seed)

	r := rand.New(rand.NewSource(seed))
	for i := 0; i < times; i++ {
		fmt.Println(r.Uint64())
	}
}
