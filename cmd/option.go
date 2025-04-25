/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

func formatOptionTable(options []Option) string {
	var rows [][]string

	for _, option := range options {
		rows = append(rows, []string{strconv.FormatUint(uint64(option.ID), 10), option.Label})
	}

	table := ltable.New().Rows(rows...).Headers("Id", "Option")
	return table.Render()
}

func formatEmptyState() string {
  style := lipgloss.NewStyle().
    Bold(true).
    PaddingTop(1).
    PaddingLeft(4).
    Foreground(lipgloss.Color("5"))

  return style.Render("There are no options in the list, use `ranker option add` to add some!")
}

// optionCmd represents the option command
var optionCmd = &cobra.Command{
	Use:     "option",
	Aliases: []string{"options"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// db := initDb(dbPath("sqlite-testing"))
		// db.Create(&Option{Label: "Test Option"})
		listData := loadLists()
		db, err := loadDb(dbPath(listData.ActiveList))
		check(err)
		options := loadOptions(db)

		if len(options) == 0 {
			fmt.Println(formatEmptyState())
		} else {
			fmt.Println("\n" + formatOptionTable(options))
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
