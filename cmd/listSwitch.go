/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listSwitchCmd represents the listSwitch command
var listSwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to another list",
	Long: `Sets the current active list to the one provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listSwitch called")
		fmt.Println("dont forget to check that the input list exists before switching to it")
	},
}

func init() {
	listCmd.AddCommand(listSwitchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listSwitchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listSwitchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
