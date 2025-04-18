/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"errors"
	"github.com/spf13/cobra"
)

// listInitCmd represents the listInit command
var listInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Start a new list of options to be ranked",
	Long: `Given a name for the list, this will create a new sqlite database in your config directory.
	The name for the list will be used as a filename.
	The new list will become the 'active' list.
	You can pass in a path to a file of options. Each unique line in the file will be stored as an option`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listInit called")
		lists := loadLists()
		listName := args[0]
		lists, err := addNewOptionList(lists, listName)
		if errors.Is(err, OptionListExists) {
			fmt.Printf("A file already exists named %s, remove it or pick a new name\n", listName)
			return
		}

		newListId := lists.ActiveList
		if fileExists(dbPath(newListId)) {
			fmt.Printf("A file already exists at %s, remove it or pick a new name\n", dbPath(newListId))
		} else {
			initDb(dbPath(newListId))
			persistListConfig(lists)
			fmt.Println("new list created & set to active")
		}

	},
}

func init() {
	listCmd.AddCommand(listInitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listInitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listInitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
