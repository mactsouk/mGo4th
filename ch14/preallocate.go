package main

import "fmt"

func main() {
	// Pre-allocate a slice with a capacity of 10
	mySlice := make([]int, 0, 100)

	// Append elements to the slice
	for i := 0; i < 100; i++ {
		mySlice = append(mySlice, i)
	}

	fmt.Println(mySlice)

	// Pre-allocate a map with an initial capacity of 10
	myMap := make(map[string]int, 10)

	// Add key-value pairs to the map
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("k%d", i)
		myMap[key] = i
	}

	fmt.Println(myMap)
}
