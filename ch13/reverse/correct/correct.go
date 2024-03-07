package main

// This version works correctly

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func R1(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Invalid UTF-8")
	}

	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	return string(a), nil
}

func R2(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Invalid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	str := "1234567890"

	R1ret, _ := R1(str)
	fmt.Println(R1ret)

	R2ret, _ := R2(str)
	fmt.Println(R2ret)
}
