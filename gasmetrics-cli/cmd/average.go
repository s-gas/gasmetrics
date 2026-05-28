package cmd

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/spf13/cobra"
)

var averageCmd = &cobra.Command{
	Use:   "average",
	Short: "Display the average consumption per day",
	Run: func(cmd *cobra.Command, args []string) {
		url := baseUrl + statsEndpoint
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		var stats Stats
		if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Daily average consumption: %.2f m³\n", stats.Avg)
	},
}

func init() {
	rootCmd.AddCommand(averageCmd)
}
