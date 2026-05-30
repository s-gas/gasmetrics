package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete reading by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var reading Reading
		var err error
		reading.Id, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if reading.Id < 1 {
			fmt.Fprintln(os.Stderr, "ID must be greater than 0")
			os.Exit(1)
		}
		url := fmt.Sprintf("%s/%d", baseUrl, reading.Id)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNoContent {
			fmt.Println("Entry deleted successfully")
		} else {
			fmt.Fprintln(os.Stderr, "Failed to delete entry: status code:", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
