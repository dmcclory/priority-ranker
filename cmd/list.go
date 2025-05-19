/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"os"

	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

func formatListTable(lists []OptionList) string {
	var rows [][]string
	activeRow := -1

	for i, list := range lists {
		created := time.Unix(int64(list.CreatedAt), 0)
		timestamp := fmt.Sprintf("%02d/%02d/%d", created.Month(), created.Day(), created.Year())
		if list.Active {
			activeRow = i
		}
		var warning string
		if !list.DbExists {
			warning = "DB is missing"
		}
		rows = append(rows, []string{list.Name, timestamp, warning})
	}

	re := lipgloss.NewRenderer(os.Stdout)
	activeStyle := re.NewStyle().Bold(true)
	inactiveStyle := re.NewStyle().Bold(false)

	// this -> https://github.com/charmbracelet/lipgloss/blob/master/table/table_test.go
	// was helpful for understanding how to construct these tables
	table := ltable.New().
		Rows(rows...).
		Headers("List", "Created On", "Warning").
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == activeRow {
				return activeStyle
			} else {
				return inactiveStyle
			}
		})

	return table.Render()
}

func emptyStateMessage() string {
	return "no lists have been created, use `ranker list init` to get started`"
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See an index of the existing lists",
	Long: `
This is a shell command that shows all of the current lists.
Use the init, delete, and switch subcommands on individual lists.`,
	Run: func(cmd *cobra.Command, args []string) {
		listData := loadLists()

		if len(listData.Lists) == 0 {
			fmt.Println(warningStyle().Render(emptyStateMessage()))
		} else {
			markListEntryAsActive(listData.ActiveList, listData)
			lists := []OptionList{}
			for _, v := range listData.Lists {
				lists = append(lists, v)
			}
			fmt.Println(formatListTable(lists))
			if listData.ActiveList == "" {
				fmt.Println(warningStyle().Render("None of these lists currently active!, use `list switch $LIST_ID` to make one active"))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
