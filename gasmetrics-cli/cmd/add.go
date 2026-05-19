/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:		"add <value>",
	Short: 	"Add a gas reading (int)",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
