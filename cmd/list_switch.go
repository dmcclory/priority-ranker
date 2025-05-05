/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// listSwitchCmd represents the listSwitch command
var listSwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to another list",
	Long:  `Sets the current active list to the one provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		lists := loadLists()

		var newId string
		options := []huh.Option[string]{}

		for _, list := range lists.Lists {
			options = append(options, huh.NewOption(list.Name, list.Id))
		}
		f := huh.NewSelect[string]().Title("Switch list:").Options(
			options...,
		).Value(&newId)

		if err := f.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
		} else {
			fmt.Println("updating the active list to: ", newId)

			updateActiveList(newId, lists)

			fmt.Println("(dont forget to check that the input list exists before switching to it)")
		}

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
