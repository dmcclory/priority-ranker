/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "See the ranking of the options",
	Long: `Prints a table of the relative ranking of the options, based on your vote history.

	It takes ~30 votes to start to get meaningful data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("results called")

		defaultStyles := table.DefaultStyles()

		styles := table.Styles {
			Cell: defaultStyles.Cell,
			Header: defaultStyles.Header,
			// Selected: defaultStyles.Cell,
		}

		table := ltable.New().Headers("Rank", "Option", "Score").
		Row("1", "Good Book", "10.4930430").
		Row("2", "Something else", "5.4930430").
		Row("3", "take a nap", "3.4930430").
		Row("4", "think about savingss", "2.4930430").
		Row("5", "cook a meal", "1.4930430").
		StyleFunc(func(row, _ int) lipgloss.Style {
			if row == 0 {
				return styles.Header
			}
			return styles.Cell
		})
		// StyleFunc(func(row, col int) lipgloss.Style {
			// switch {
			// case row == 0:
				// return Hea
		// })

		// t := table.New(
			// table.WithColumns(columns),
			// table.WithRows(rows),
			// table.WithFocused(true),
			// table.WithHeight(10),
		// )


		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resultsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resultsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
