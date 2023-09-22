/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// oneCmd represents the one command
var oneCmd = &cobra.Command{
	Use:     "one",
	Aliases: []string{"cmd1"},
	Short:   "Command one",
	Long: `A longer description
				for command one.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("one called")
	},
}

func init() {
	rootCmd.AddCommand(oneCmd)
}
