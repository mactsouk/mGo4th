package main

// https://go101.org/article/unsafe.html

import (
	"fmt"
	"unsafe"
)

func byteToString(bStr []byte) string {
	if len(bStr) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(bStr), len(bStr))
}

func stringToByte(str string) []byte {
	if str == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func main() {
	str := "Go!"
	d := unsafe.StringData(str)
	b := unsafe.Slice(d, len(str))
	// byte is an alias for uint8
	fmt.Printf("Type %T contains %s\n", b, b)

	sData := []int{10, 20, 30, 40}
	// Get the memory address of sData
	fmt.Println("Pointer:", unsafe.SliceData(sData))

	// String to Byte slice
	var hi string = "Mastering Go, 4th edition!"
	myByteSlice := stringToByte(hi)
	fmt.Printf("myByteSlice type: %T\n", myByteSlice)

	// Byte slice to string
	myStr := byteToString(myByteSlice)
	fmt.Printf("myStr type: %T\n", myStr)
}
