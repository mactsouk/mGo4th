package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	var initialized = 0

	nValues := 0
	var sum float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		nValues = nValues + 1
		sum = sum + n

		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Number of values:", nValues)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// Mean value
	if nValues == 0 {
		return
	}
	meanValue := sum / float64(nValues)
	fmt.Printf("Mean value: %.5f\n", meanValue)

	// Standard deviation
	// https://en.wikipedia.org/wiki/Standard_deviation
	var squared float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		squared = squared + math.Pow((n-meanValue), 2)
	}

	standardDeviation := math.Sqrt(squared / float64(nValues))
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)
}
