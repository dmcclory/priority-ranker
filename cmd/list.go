/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)


type ChoiceList struct {
	Name string
	Id string
	Active bool
}

func getLists() []ChoiceList {
	return []ChoiceList{
		{Id: "first-list", Name: "First List"},
		{Id: "second-list", Name: "Second List"},
		{Id: "third-list", Name: "Third List"},
	}
}

func getEmptyListResult() []ChoiceList {
	return []ChoiceList{}
}

func getListNames(lists []ChoiceList) []string {
	var names []string
	for _, list := range lists {
		names = append(names, list.Name)
	}
	return names
}

func markActiveListEntry(lists []ChoiceList, activeList string) []ChoiceList {
	for i := range lists {
		if lists[i].Id == activeList {
			lists[i].Active = true
		}
	}
	return lists
}

func formatTable(lists []ChoiceList) string {
	var rows [][]string

	for _, list := range lists {
		rows = append(rows, []string{list.Name, strconv.FormatBool(list.Active)})
	}

	// this -> https://github.com/charmbracelet/lipgloss/blob/master/table/table_test.go
	// was helpful for understanding how to construct these tables
	table := ltable.New().Rows(rows...).Headers("List", "Active")

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
		lists := getLists()

		if len(lists) == 0 {
		  fmt.Println(emptyStateMessage())
		} else {
			// gotta figure out how to test this method
			activeList := loadActiveProjectFromFile()
			lists = markActiveListEntry(lists, activeList)
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
