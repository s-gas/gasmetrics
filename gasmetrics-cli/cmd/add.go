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
		body := strings.NewReader(fmt.Sprintf(`{"value": %s}`, args[0]))
		resp, err := http.Post(url, contentType, body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode == http.StatusOK {
			log.Println("Entry added successfully")
		} else {
			log.Println("Failed to add entry: status code:", resp.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
