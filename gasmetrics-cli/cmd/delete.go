package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
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
			log.Fatal(err)
		}
		if reading.Id < 1 {
			log.Fatal("ID must be greater than 0")
		}
		url := fmt.Sprintf("%s/%d", baseUrl, reading.Id)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNoContent {
			log.Println("Entry deleted successfully")
		} else {
			log.Println("Failed to delete entry: status code:", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
