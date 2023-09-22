/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log/slog"
	"math"
	"os"
	"slices"
	"strconv"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert command",
	Long: `The insert command reads a datafile and stores
	its data into the application in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
		// Work with logger
		if disableLogging == false {
			logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
		}

		slog.SetDefault(logger)

		if file == "" {
			logger.Info("Need a file to read!")
			return
		}

		_, ok := index[file]
		if ok {
			fmt.Println("Found key:", file)
			delete(index, file)
		}

		// Now, delete it from data
		if ok {
			for i, k := range data {
				if k.Filename == file {
					data = slices.Delete(data, i, i+1)
					break
				}
			}
		}

		err := ProcessFile(file)
		if err != nil {
			s := fmt.Sprintf("Error processing: %s", err)
			logger.Warn(s)
		}

		err = saveJSONFile(JSONFILE)
		if err != nil {
			s := fmt.Sprintf("Error saving data: %s", err)
			logger.Info(s)
		}
	},
}

var file string

func init() {
	rootCmd.AddCommand(insertCmd)
	// define required local flag
	insertCmd.Flags().StringVarP(&file, "file", "f", "", "Filename to process")
	insertCmd.MarkFlagRequired("file")

	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// Work with logger
	if disableLogging == false {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}

	slog.SetDefault(logger)
	s := fmt.Sprintf("%d records in total.", len(data))
	logger.Info(s)
}

func readFile(filepath string) ([]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	for _, line := range lines {
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			fmt.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}

	return values, nil
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

func ProcessFile(file string) error {
	currentFile := Entry{}
	currentFile.Filename = file

	values, err := readFile(file)
	if err != nil {
		return err
	}

	currentFile.Len = len(values)
	currentFile.Minimum = slices.Min(values)
	currentFile.Maximum = slices.Max(values)
	meanValue, standardDeviation := stdDev(values)
	currentFile.Mean = meanValue
	currentFile.StdDev = standardDeviation

	data = append(data, currentFile)

	return nil
}
