/*
Copyright © 2023 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of delete",
	Long: `A longer description
	for the delete command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	threeCmd.AddCommand(deleteCmd)
}
