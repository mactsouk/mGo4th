package exampleFunctions

import "fmt"

func ExampleLengthRange() {
	fmt.Println(LengthRange("Mihalis"))
	fmt.Println(LengthRange("Mastering Go, 4th edition!"))
	// Output:
	// 7
	// 7
}
