package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the time from the RESTful server",
	Long:  `This command mainly exists for making sure that the server works.`,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("GET", SERVER+PORT+"/time", nil)
		if err != nil {
			fmt.Println("Timefunction – Error in req: ", err)
			return
		}

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp == nil || (resp.StatusCode == http.StatusNotFound) {
			fmt.Println(resp)
			return
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		fmt.Print(string(data))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
