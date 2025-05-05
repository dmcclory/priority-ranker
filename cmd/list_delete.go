/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listDeleteCmd represents the listDelete command
var listDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a list",
	Long:  `Deletes the list. Be careful! This operation is irreversible`,
	Run: func(cmd *cobra.Command, args []string) {
		lists := loadLists()
		listId := args[0]
		_, present := lists.Lists[listId]

		if !present {
			fmt.Printf("%s is not a list, exiting\n", listId)
			return
		}

		lists, err := deleteList(lists, listId)
		check(err)
		if err != nil {
			fmt.Printf("error while deleting list: %s\n", listId )
		} else {
		  persistListConfig(lists)
			fmt.Printf("successfully deleted %s", listId)
		}
	},
}

func init() {
	listCmd.AddCommand(listDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
