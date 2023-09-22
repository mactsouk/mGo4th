/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	sort.Sort(DFslice(data))
	text, err := PrettyPrintJSONstream(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)

	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// Work with logger
	if disableLogging == false {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}

	slog.SetDefault(logger)
	s := fmt.Sprintf("%d records in total.", len(data))
	logger.Info(s)
}

// PrettyPrintJSONstream pretty prints the contents of the phone book
func PrettyPrintJSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// Implement sort.Interface
func (a DFslice) Len() int {
	return len(a)
}

func (a DFslice) Less(i, j int) bool {
	if a[i].Mean == a[j].Mean {
		return a[i].StdDev < a[j].StdDev
	}
	return a[i].Mean < a[j].Mean
}

func (a DFslice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
