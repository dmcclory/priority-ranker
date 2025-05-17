/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.2"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ranker",
	Version: version,
	Short: "a tool for finding the relative priority in a group of options",
	Long: `
ranker is a tool for finding the priority in a list of options.
You can create and switch between lists with the 'list init' and 'list subcommands'

'ranker vote' - You find the relative priority by voting on randomly drawn pairs of options.

'ranker results' - Once you've voted a number of times, you can see a ranking.

The rankings are calculated using the Bradley Terry model.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
