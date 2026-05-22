package cmd

import (
		"time"
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
		reading := parseReading(cmd, args)
		body, err := json.Marshal(reading)
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

func parseReading(cmd *cobra.Command, args []string) Reading {
		var reading Reading
		var err error
		dateStr, _ := cmd.Flags().GetString("date")
		if !cmd.Flags().Changed("date") {
			reading.Date = time.Now()	
		} else {
			reading.Date, err = time.Parse("2006-01-02", dateStr)
			if err != nil {
				log.Fatal("invalid date format")
			}
		}
		reading.Value, err = strconv.Atoi(args[0])
		if err != nil || reading.Value <= 0 {
			log.Fatal("value must be a positive integer")
		}
		return reading

}

func init() {
	addCmd.Flags().StringP("date", "d", "", "Specific date in format YYYY-MM-DD")
	rootCmd.AddCommand(addCmd)
}
