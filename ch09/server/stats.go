package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"slices"
	"time"
)

type Entry struct {
	Name    string
	Len     int
	Minimum float64
	Maximum float64
	Mean    float64
	StdDev  float64
}

func process(file string, values []float64) Entry {
	currentEntry := Entry{}
	currentEntry.Name = file

	currentEntry.Len = len(values)
	currentEntry.Minimum = slices.Min(values)
	currentEntry.Maximum = slices.Max(values)
	meanValue, standardDeviation := stdDev(values)
	currentEntry.Mean = meanValue
	currentEntry.StdDev = standardDeviation

	return currentEntry
}

func stdDev(x []float64) (float64, float64) {
	sum := float64(0)
	for _, val := range x {
		sum = sum + val
	}

	meanValue := sum / float64(len(x))

	// Standard deviation
	var squared float64
	for i := 0; i < len(x); i++ {
		squared = squared + math.Pow((x[i]-meanValue), 2)
	}

	standardDeviation := math.Sqrt(squared / float64(len(x)))
	return meanValue, standardDeviation
}

// JSONFILE resides in the current directory
var JSONFILE = "./data.json"

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	err = Serialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
			return nil
		}
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = DeSerialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Name
		index[key] = i
	}
}

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Name]
	if ok {
		return fmt.Errorf("%s already exists", pS.Name)
	}

	data = append(data, *pS)
	// Update the index
	createIndex()

	err := saveJSONFile(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := saveJSONFile(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	return &data[i]
}

func list() string {
	var all string

	for _, k := range data {
		all = all + fmt.Sprintf("%s\t%d\t%f\t%f\n", k.Name, k.Len, k.Mean, k.StdDev)
	}
	return all
}

func main() {
	err := readJSONFile(JSONFILE)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		return
	}
	createIndex()

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(listHandler))
	mux.Handle("/insert/", http.HandlerFunc(insertHandler))
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/search", http.HandlerFunc(searchHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
