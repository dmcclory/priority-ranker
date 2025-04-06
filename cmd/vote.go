/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var choice = ""

// voteCmd represents the vote command
var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vote called")
		huh.NewSelect[string]().
		  Title("Which would you like to do first?").
			Options(
				huh.NewOption("Cool Choice 1", "cool-choice"),
				huh.NewOption("Cool Choice 2", "cool-choice-2"),
			).
			Value(&choice).Run()
			fmt.Println("vote result: ", choice)
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
