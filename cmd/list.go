/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)


func formatTable(lists []OptionList) string {
	var rows [][]string

	for _, list := range lists {
		created := time.Unix(int64(list.CreatedAt), 0)
		timestamp := fmt.Sprintf("%d/%d/%d", created.Year(), created.Month(), created.Day())
		var warning string
    if !list.DbExists {
			warning = "DB is missing"
		}
		rows = append(rows, []string{list.Name, strconv.FormatBool(list.Active), timestamp, warning})
	}

	// this -> https://github.com/charmbracelet/lipgloss/blob/master/table/table_test.go
	// was helpful for understanding how to construct these tables
	table := ltable.New().Rows(rows...).Headers("List", "Active", "Created On", "Warning")

	return table.Render()
}

func emptyStateMessage() string {
  return "no lists have been created, use `ranker list init` to get started`"
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See an index of the existing lists",
	Long: `This is a shell command that shows all of the current lists.
	Use the init, delete, and switch subcommands on individual lists.`,
	Run: func(cmd *cobra.Command, args []string) {
		listData := loadLists()

		if len(listData.Lists) == 0 {
		  fmt.Println(emptyStateMessage())
		} else {
			// gotta figure out how to test this method
			// lists := markListEntryAsActiveMap(listData, listData.ActiveList)
			markListEntryAsActive(listData.ActiveList, listData)
			lists := []OptionList{}
			for _, v := range listData.Lists {
				lists = append(lists, v)
			}
			fmt.Println(formatTable(lists))
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
