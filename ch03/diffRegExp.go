package main

// https://stackoverflow.com/questions/25159236/panic-when-compiling-a-regular-expression

import (
	"fmt"
	"regexp"
)

func main() {
	// This is a raw string literal
	var re string = `^.*(?=.{7,})(?=.*\d)$`

	exp1, err := regexp.Compile(re)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("RegExp:", exp1)

	exp2 := regexp.MustCompile(re)
	fmt.Println(exp2)
}
