package main

import (
	"fmt"
	"os"

	"github.com/mactsouk/sqlite06"
)

func main() {
	sqlite06.Filename = "ch06.db"

	db, err := sqlite06.OpenConnection()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Connection string:", db)
	os.Remove(sqlite06.Filename)
}
