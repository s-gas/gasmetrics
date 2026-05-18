package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:		"add <value>",
	Short: 	"Add a gas reading",
	Args:		cobra.ExactArgs(1),
	Run:		func(cmd *cobra.Command, args []string) {
		url := "http://hauslab/readings"
		contentType := "application/json"
		body := strings.NewReader(fmt.Sprintf(`{"value": %s}`, args[0]))
		resp, err := http.Post(url, contentType, body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode == http.StatusOK {
			log.Println("Entry added successfully")
		} else {
			log.Println("Failed to add entry")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
