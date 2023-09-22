/*
Copyright © 2021 Mihalis Tsoukalos <mihalistsoukalos@gmail.com>
*/

package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert new entries",
	Long: `This command inserts new data to the
	Phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		dataset, _ := cmd.Flags().GetString("dataset")
		if dataset == "" {
			fmt.Println("Dataset is empty!")
			return
		}

		values, _ := cmd.Flags().GetString("values")
		if values == "" {
			fmt.Println("No data!")
			return
		}

		VALS := strings.Split(values, ",")
		vSend := ""
		for _, v := range VALS {
			_, err := strconv.ParseFloat(v, 64)
			if err == nil {
				vSend = vSend + "/" + v
			}
		}

		// Create request in two steps for readability
		URL := "http://" + SERVER + ":" + PORT + "/insert/"
		URL = URL + "/" + dataset + "/" + vSend + "/"

		// Send request to server
		data, err := http.Get(URL)
		if err != nil {
			fmt.Println("**", err)
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
			fmt.Println("*", err)
			return
		}

		fmt.Print(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP("dataset", "d", "", "Dataset name")
	insertCmd.Flags().StringP("values", "v", "", "List of values")
}
