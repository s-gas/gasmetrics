package cmd

import (
		"os"
		"time"
		"fmt"
		"bytes"
		"encoding/json"
		"strconv"
    "net/http"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:		"add <value>",
	Short: 	"Add a gas reading in m³ (int)",
	Args:		cobra.ExactArgs(1),
	Run:		func(cmd *cobra.Command, args []string) {
		reading, err := parseReading(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		body, err := json.Marshal(reading)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		resp, err := http.Post(baseUrl, contentType, bytes.NewReader(body))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode == http.StatusCreated {
			fmt.Println("Entry added successfully")
		} else {
			fmt.Fprintln(os.Stderr, "Failed to add entry: status code:", resp.StatusCode)
		}
	},
}

func parseReading(cmd *cobra.Command, args []string) (Reading, error) {
		var reading Reading
		var err error
		dateStr, _ := cmd.Flags().GetString("date")
		if !cmd.Flags().Changed("date") {
			reading.Date = time.Now()	
		} else {
			reading.Date, err = time.Parse("2006-01-02", dateStr)
			if err != nil {
				return Reading{}, fmt.Errorf("parseReading: invalid date format")
			}
			if reading.Date.After(time.Now()) {
				return Reading{}, fmt.Errorf("parseReading: invalid date")
			}
		}
		reading.Value, err = strconv.Atoi(args[0])
		if err != nil || reading.Value <= 0 {
			return Reading{}, fmt.Errorf("parseReading: value must be an integer")
		}
		return reading, nil
}

func init() {
	addCmd.Flags().StringP("date", "d", "", "Specific date in format YYYY-MM-DD")
	rootCmd.AddCommand(addCmd)
}
