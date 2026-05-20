package cmd

import (
    "log"
		"bytes"
		"encoding/json"
		"strconv"
    "net/http"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:		"add <value>",
	Short: 	"Add a gas reading (int)",
	Args:		cobra.ExactArgs(1),
	Run:		func(cmd *cobra.Command, args []string) {
		value, err := strconv.Atoi(args[0])
		if err != nil || value <= 0 {
			log.Fatal("value must be a positive integer")
		}
		body, err := json.Marshal(Reading{Value: value})
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Post(baseUrl, contentType, bytes.NewReader(body))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode == http.StatusCreated {
			log.Println("Entry added successfully")
		} else {
			log.Println("Failed to add entry: status code:", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
