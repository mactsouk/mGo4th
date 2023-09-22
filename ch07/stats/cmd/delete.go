/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"slices"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete command",
	Long:  `A command for deleting data.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
		slog.SetDefault(logger)

		_, ok := index[key]
		if ok {
			logger.Info("Found key:", key)
			fmt.Println("Found key:", key)
			delete(index, key)
		} else {
			s := fmt.Sprintf("%s not found!", key)
			logger.Info(s)
			return
		}

		// Now, delete it from data
		fmt.Println(data)
		for i, k := range data {
			if k.Filename == key {
				data = slices.Delete(data, i, i+1)
				break
			}
		}

		err := saveJSONFile(JSONFILE)
		if err != nil {
			logger.Warn("Error saving data:", err)
		}

		s := fmt.Sprintf("Deleting key %s:", key)
		logger.Info(s)
	},
}

var key string

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key to delete")
	deleteCmd.MarkFlagRequired("key")
}
