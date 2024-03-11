package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	times := 2
	seed := int64(0)
	arguments := os.Args
	if len(arguments) > 1 {
		n, err := strconv.ParseInt(arguments[1], 10, 64)
		if err == nil {
			seed = n
		}
	}

	fmt.Println("Using seed:", seed)

	src := rand.NewSource(seed)
	r := rand.New(src)
	for i := 0; i < times; i++ {
		fmt.Println(r.Uint64())
	}
}
