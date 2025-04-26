/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var optionsInput bool
var optionsFile string

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
		listName := strings.Join(args, " ")
		lists, err := addNewOptionList(lists, listName)
		if errors.Is(err, OptionListExists) {
			fmt.Printf("A file already exists named %s, remove it or pick a new name\n", listName)
			return
		}

		newListId := lists.ActiveList
		if fileExists(dbPath(newListId)) {
			fmt.Printf("A file already exists at %s, remove it or pick a new name\n", dbPath(newListId))
			return
		}

		db, err := initDb(dbPath(newListId))
		persistListConfig(lists)
		fmt.Println("new list created & set to active")

		optionLabels := []string{}

		if optionsInput {
			var textInput string
			huh.NewText().
			  Title(fmt.Sprintf("Enter options, each line will become part of the '%s' list\n(Ctrl+J to add a new line)", listName)).
				Lines(10).
				ShowLineNumbers(true).
				Value(&textInput).
				Run()
				optionLabels = strings.Split(textInput, "\n")
		}
		if optionsFile != "" {
			data, err := os.ReadFile(optionsFile)
			if err != nil {
				fmt.Printf("not able to read %s, the options have not been added to the db", optionsFile)
			} else {
				lines := strings.Split(string(data), "\n")

				if len(lines) > 0 {
					optionLabels = append(optionLabels, lines...)
				}
			}
		}

		// handle duplicate entries as well
		noBlanks := []string{}
		for _, line := range optionLabels {
			if line != "" {
				noBlanks = append(noBlanks, line)
			}
		}

		if len(noBlanks) > 0 {
			_, err = addOptions(db, noBlanks)
				fmt.Println("not able to add options to the database", err)
			if err != nil {
			}
			fmt.Printf("added %d options to the list\n", len(noBlanks))
		}

	},
}

func init() {
	listCmd.AddCommand(listInitCmd)
	listInitCmd.Flags().BoolVarP(&optionsInput, "options", "o", false, "add options with a text input")
	listInitCmd.Flags().StringVarP(&optionsFile, "options-file", "f", "", "read options from a file")
}
