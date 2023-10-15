package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	arguments := os.Args[1:]

	for i, arg := range arguments {
		t, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		switch i {
		case 0:
			MIN = t
			MAX = MIN + 100
		case 1:
			MAX = t
		case 2:
			TOTAL = t
		case 3:
			SEED = int64(t)
		}
	}

	switch len(arguments) {
	case 1, 2, 3, 4:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}
