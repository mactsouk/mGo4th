/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all entries.",
	Long:  `List the contents of Phone book.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		// Create request
		URL := "http://" + SERVER + ":" + PORT + "/list"

		// Send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Check HTTP Status Code
		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		// Read data
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("List of entries:")
		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
