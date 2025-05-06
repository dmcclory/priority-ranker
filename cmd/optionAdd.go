/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// optionAddCmd represents the optionAdd command
var optionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new option to the current list",
	Long: `
Adds a new option to the current list of options. If the input text is an exact match of an already existing option, it will be skipped.`,
	Run: func(cmd *cobra.Command, args []string) {
		option := strings.Join(args, " ")

		listData := loadLists()
		
		db, err := loadDb(dbPath(listData.ActiveList))
		check(err)

		// how do we handle errors for insert and query?
		addOption(db, option)

		fmt.Printf("added '%s' to the active list of options\n", option)
	},
}

func init() {
	optionCmd.AddCommand(optionAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
