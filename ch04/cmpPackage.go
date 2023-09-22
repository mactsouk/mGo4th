package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(cmp.Compare(5, 4))
	fmt.Println(cmp.Compare(4, 5))
	fmt.Println(cmp.Less(4, 5))
}
