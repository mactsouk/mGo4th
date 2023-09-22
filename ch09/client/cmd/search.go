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

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a telephone number",
	Long:  `This command searches for a given phone number..`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		dataset, _ := cmd.Flags().GetString("dataset")
		if dataset == "" {
			fmt.Println("Dataset name is empty!")
			return
		}

		// Create request
		URL := "http://" + SERVER + ":" + PORT + "/search/" + dataset

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

		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("dataset", "d", "", "Dataset name to search")
}
