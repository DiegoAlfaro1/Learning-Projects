package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "OpenGrep CLI Tool",
	Short: "This is a CLI tool to use opengrep easily",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error ocurred while executing the tool '%s'\n", err)
		os.Exit(1)
	}
}
