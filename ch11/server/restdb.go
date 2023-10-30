package main

// Database: PostgreSQL
//
// Functions to support the interaction with the database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Filename = "REST.db"

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", Filename)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return nil
	}
	return db
}

// FromJSON decodes a serialized JSON record - User{}
func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON encodes a User JSON record
func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// DeleteUser is for deleting a user defined by ID
func DeleteUser(ID int) bool {
	db := OpenConnection()
	if db == nil {
		log.Println("Cannot connect to SQLite3!")
		db.Close()
		return false
	}
	defer db.Close()

	// Check is the user ID exists
	t := FindUserID(ID)
	if t.ID == 0 {
		log.Println("User", ID, "does not exist.")
		return false
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE UserID = $1")
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}

	return true
}

// InsertUser is for adding a new user to the database
func InsertUser(u User) bool {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		return false
	}
	defer db.Close()

	if IsUserValid(u) {
		log.Println("User", u.Username, "already exists!")
		return false
	}

	stmt, err := db.Prepare("INSERT INTO users(username, password, lastlogin, admin, active) values($1,$2,$3,$4,$5)")
	if err != nil {
		log.Println("Adduser:", err)
		return false
	}

	stmt.Exec(u.Username, u.Password, u.LastLogin, u.Admin, u.Active)
	return true
}

// ListAllUsers is for returning all users from the database table
func ListAllUsers() []User {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		db.Close()
		return []User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users \n")
	if err != nil {
		log.Println(err)
		return []User{}
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}

	log.Println("All:", all)
	return all
}

// ListLogged is for returning all logged users
// This was created by mistake - the server uses
// ReturnLoggedUsers() instead!
func ListLogged() []User {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite!")
		db.Close()
		return []User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE active = 1 \n")
	if err != nil {
		log.Println(err)
		return []User{}
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		_ = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}

	log.Println("All:", all)
	return all
}

// FindUserID is for returning a user record defined by ID
func FindUserID(ID int) User {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite!")
		db.Close()
		return User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where UserID = $1 \n", ID)
	if err != nil {
		log.Println("Query:", err)
		return User{}
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}
		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}
	return u
}

// FindUserUsername is for returning a user record defined by a username
func FindUserUsername(username string) User {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		db.Close()
		return User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where username = $1 \n", username)
	if err != nil {
		log.Println("FindUserUsername Query:", err)
		return User{}
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}
		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}
	return u
}

// ReturnLoggedUsers is for returning all logged in users
func ReturnLoggedUsers() []User {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		db.Close()
		return []User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE active = 1 \n")
	if err != nil {
		log.Println(err)
		return []User{}
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return []User{}
		}
		temp := User{c1, c2, c3, c4, c5, c6}
		log.Println("temp:", all)
		all = append(all, temp)
	}

	log.Println("Logged in:", all)
	return all
}

// IsUserAdmin determines whether a user is
// an administrator or not
func IsUserAdmin(u User) bool {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		db.Close()
		return false
	}
	defer db.Close()

	statement := fmt.Sprintf(`SELECT * FROM users WHERE username = '%s'`, u.Username)
	rows, err := db.Query(statement)
	if err != nil {
		log.Println("IsUserAdmin:", err)
		return false
	}

	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	// If there exist multiple users with the same username,
	// we will get the FIRST ONE only.
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println("IsUserAdmin:", err)
			return false
		}
		temp = User{c1, c2, c3, c4, c5, c6}
	}

	if u.Username == temp.Username && u.Password == temp.Password && temp.Admin == 1 {
		return true
	}
	return false
}

func IsUserValid(u User) bool {
	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE username = $1 \n", u.Username)
	if err != nil {
		log.Println(err)
		return false
	}

	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	// If there exist multiple users with the same username,
	// we will get the FIRST ONE only.
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return false
		}
		temp = User{c1, c2, c3, c4, c5, c6}
	}

	if u.Username == temp.Username && u.Password == temp.Password {
		return true
	}
	return false
}

// UpdateUser allows you to update user name
func UpdateUser(u User) bool {
	log.Println("Updating user:", u)

	db := OpenConnection()
	if db == nil {
		fmt.Println("Cannot connect to SQLite3!")
		db.Close()
		return false
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET username=$1, password=$2, admin=$3, active=$4 WHERE UserID = $5")
	if err != nil {
		log.Println("Adduser:", err)
		return false
	}

	res, err := stmt.Exec(u.Username, u.Password, u.Admin, u.Active, u.ID)
	if err != nil {
		log.Println("UpdateUser failed:", err)
		return false
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Println("RowsAffected() failed:", err)
		return false
	}
	log.Println("Affected:", affect)
	return true
}
