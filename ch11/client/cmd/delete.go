package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deleting users",
	Long:  `This command deletes existing users from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/username"
		user := User{Username: username, Password: password}

		// Convert data string to User Structure
		var u2 User
		err := json.Unmarshal([]byte(data), &u2)
		if err != nil {
			fmt.Println("Unmarshal:", err)
			return
		}

		// bytes.Buffer is both a Reader and a Writer
		buf := new(bytes.Buffer)
		err = user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		URL := SERVER + PORT + endpoint + "/" + fmt.Sprint(u2.ID)
		req, err := http.NewRequest(http.MethodDelete, URL, buf)
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
			fmt.Println("User with ID", u2.ID, "deleted.")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
