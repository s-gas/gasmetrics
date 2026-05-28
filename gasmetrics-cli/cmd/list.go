package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List readings",
	Run: func(cmd *cobra.Command, args []string) {
		limit, err := cmd.Flags().GetUint("limit")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if limit == 0 {
			return
		}
		url := fmt.Sprintf("%s?limit=%d", baseUrl, limit)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		fmt.Println("ID\tVALUE\tDATE")
		var readings []Reading
		if err := json.NewDecoder(resp.Body).Decode(&readings); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, reading := range readings {
			fmt.Printf("%v\t%vm³\t%v\n", reading.Id, reading.Value, reading.Date.Format("2006-01-02"))
		}
	},
}

func init() {
	listCmd.Flags().UintP("limit", "l", 10, "Number of readings to list")
	rootCmd.AddCommand(listCmd)
}
