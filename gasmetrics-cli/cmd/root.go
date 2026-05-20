package cmd

import (
	"os"
	"time"
	"github.com/spf13/cobra"
)

const (
	url = "http://localhost:1024/readings"
	// url = "http://hauslab/readings"
	contentType = "application/json"
)

type Reading struct {
	Value int `json:"value"`
	Date time.Time `json:"recorded_at"`
}

var rootCmd = &cobra.Command{
	Use:   "gasmetrics-cli",
	Short: "CLI tool to submit gas readings",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


