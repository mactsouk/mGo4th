package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const PORT = ":1234"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get dataset
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)
	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)

	dataset := paramStr[2]
	err := deleteEntry(dataset)
	if err != nil {
		fmt.Println(err)
		Body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", Body)
		return
	}

	Body := dataset + " deleted!\n"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", Body)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := list()
	fmt.Fprintf(w, "%s", Body)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := fmt.Sprintf("Total entries: %d\n", len(data))
	fmt.Fprintf(w, "%s", Body)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Not enough arguments: "+r.URL.Path)
		return
	}

	dataset := paramStr[2]

	// These are string values
	dataStr := paramStr[3:]
	data := make([]float64, 0)

	for _, v := range dataStr {
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			data = append(data, val)
		}
	}

	entry := process(dataset, data)
	err := insert(&entry)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		Body := "Failed to add record\n"
		fmt.Fprintf(w, "%s", Body)
	} else {
		Body := "New record added successfully\n"
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", Body)
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Get Search value from URL
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}

	var Body string
	dataset := paramStr[2]
	t := search(dataset)
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		Body = "Could not be found: " + dataset + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		Body = fmt.Sprintf("%s %d %f %f\n", t.Name, t.Len, t.Mean, t.StdDev)
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)
	fmt.Fprintf(w, "%s", Body)
}
