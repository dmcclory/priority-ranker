/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCurrentCmd represents the listCurrent command
var listCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show information about the current active list",
	Long:  `
Prints out data about the current list of options which you are ranking.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listCurrent called")
		fmt.Println("it would be nice to give this a little table")
		fmt.Println("it could have - name, date created, number of options, number of votes, date of the last vote")
	},
}

func init() {
	listCmd.AddCommand(listCurrentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCurrentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCurrentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
