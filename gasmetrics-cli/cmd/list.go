package cmd

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List readings",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var readings []Reading
		if err := json.NewDecoder(resp.Body).Decode(&readings); err != nil {
			log.Fatal(err)
		}
		for _, reading := range readings {
			fmt.Printf("%v: %v m³\n", reading.Date.Format("2006-01-02"), reading.Value) 
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
