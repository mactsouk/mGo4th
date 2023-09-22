package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect or Create an SQLite database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer db.Close()

	// Get SQLite3 version
	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		fmt.Println("Version:", err)
		return
	}

	fmt.Println("SQLite3 version:", version)
	os.Remove("test.db")
}
