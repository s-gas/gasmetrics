package cmd

import (
	"log"
	"time"
	"github.com/spf13/cobra"
)

const (
	baseUrl = "http://hauslab/readings"
	statsEndpoint = "/stats"
	contentType = "application/json"
)

type Reading struct {
	Id 		int 			`json:"id"`
	Value int 			`json:"value"`
	Date 	time.Time `json:"recorded_at"`
}

type Stats struct {
	Avg		float64		`json:"avg"`
}

var rootCmd = &cobra.Command{
	Use:   "gasmetrics-cli",
	Short: "CLI tool to submit gas readings",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


