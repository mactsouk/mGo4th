/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twoCmd represents the two command
var twoCmd = &cobra.Command{
	Use:     "two",
	Aliases: []string{"cmd2"},
	Short:   "Command two",
	Long: `A longer description
	for command two.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println("Error reading username:", err)
		} else {
			fmt.Println("Username:", username)
		}
	},
}

func init() {
	rootCmd.AddCommand(twoCmd)

	twoCmd.Flags().StringP("username", "u", "Mike", "Username")
}
