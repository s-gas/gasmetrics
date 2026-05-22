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
		limit, _ := cmd.Flags().GetUint("limit")
		url := fmt.Sprintf("%s?limit=%d", baseUrl, limit)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Println("ID\tVALUE\tDATE")
		var readings []Reading
		if err := json.NewDecoder(resp.Body).Decode(&readings); err != nil {
			log.Fatal(err)
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
