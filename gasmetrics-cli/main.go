package main

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "gasmetrics-cli"}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

