package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Println("Enabling logging!")
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(os.Stderr)
		log.Println("Disabling logging!")
		log.SetOutput(io.Discard)
		log.Println("NOT GOING TO GET THAT!")
	}
}
