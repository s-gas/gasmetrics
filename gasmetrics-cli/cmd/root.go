package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	url = "http://hauslab/readings"
	contentType = "application/json"
)

type Reading struct {
	Value int `json:"value"`
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


