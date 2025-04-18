/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// optionCmd represents the option command
var optionCmd = &cobra.Command{
	Use:   "option",
	Short: "A brief description of your command",
	Aliases: []string{"options"},
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("option called")
	 	// db := initDb(dbPath("sqlite-testing"))
	  // db.Create(&Option{Label: "Test Option"})
	 	db := loadDb(dbPath("sqlite-testing"))
		options := loadOptions(db)
		for _, option := range options {
			fmt.Println("option: ", option.Label)
		}
	},
}

func init() {
	rootCmd.AddCommand(optionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
