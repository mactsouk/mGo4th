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

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Prints the status of the server.",
	Long: `This command prints information about the
	status of the phone book server.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		// Create request
		URL := "http://" + SERVER + ":" + PORT + "/status"

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
	rootCmd.AddCommand(statusCmd)
}
