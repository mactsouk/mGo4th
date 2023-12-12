package main

import (
	"fmt"
	"golang.org/x/text/language"
)

func main() {
	greece := language.Make("el")
	en := language.Make("en")
	fmt.Println(greece.Region())
	fmt.Println(en.Region())
}
