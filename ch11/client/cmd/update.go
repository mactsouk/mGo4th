package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update user data",
	Long:  `Update the user data.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/update"

		// User issuing the command
		u1 := User{Username: username, Password: password}

		// Convert data string to User Structure
		// User to be updated
		var u2 User
		err := json.Unmarshal([]byte(data), &u2)
		if err != nil {
			fmt.Println("Unmarshal:", err)
			return
		}

		users := []User{}
		users = append(users, u1)
		users = append(users, u2)

		// bytes.Buffer is both a Reader and a Writer
		buf := new(bytes.Buffer)
		err = SliceToJSON(users, buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodPut, SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("GetAll – Error in req: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("Do:", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Status code:", resp.Status)
		} else {
			fmt.Println("User", u2.Username, "updated.")
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
