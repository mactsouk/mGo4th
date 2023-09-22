/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// threeCmd represents the three command
var threeCmd = &cobra.Command{
	Use:     "three",
	Aliases: []string{"cmd3"},
	Short:   "Command three",
	Long: `A longer description
	for command three.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("three called")
	},
}

func init() {
	rootCmd.AddCommand(threeCmd)
}
