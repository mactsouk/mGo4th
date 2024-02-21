package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// CSV file read all at once
	// lines is a [][]string variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	csvwriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		err = csvwriter.Write(temp)
		if err != nil {
			return err
		}
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Println("csvData input output!")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := readCSVFile(input)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// CSV data is read in columns - each line is a slice
	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		log.Println(temp)
	}

	err = saveCSVFile(output)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
