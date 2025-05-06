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

func getListInitOptions(huhInput string, fileInput string) []string {
	var optionLabels []string
  optionLabels = strings.Split(huhInput, "\n")
	optionLabels = append(optionLabels, strings.Split(fileInput, "\n")...)

	noBlanks := []string{}
	for _, line := range optionLabels {
		if line != "" {
			noBlanks = append(noBlanks, line)
		}
	}

	optionsMap := map[string]bool{}

	for _, option := range noBlanks {
		if !optionsMap[option] {
			optionsMap[option] = true
		}
	}

	keys := make([]string, len(optionsMap))
	i := 0
	for k := range optionsMap {
		keys[i] = k
		i++
	}

	return keys
}

// listInitCmd represents the listInit command
var listInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Start a new list of options to be ranked",
	Long: `
Given a name for the list, this will create a new sqlite database in your config directory.
The name for the list will be used as a filename.
The new list will become the 'active' list.
You can pass in a path to a file of options. Each unique line in the file will be stored as an option`,
	Run: func(cmd *cobra.Command, args []string) {
		lists := loadLists()
		listName := strings.Join(args, " ")
		if listName == ""{
			fmt.Printf("No list name given, exiting\n")
			return
		}
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
		var textInput string
		var fileInput string

		if optionsInput {
			huh.NewText().
			  Title(fmt.Sprintf("Enter options, each line will become part of the '%s' list\n(Ctrl+J to add a new line)", listName)).
				Lines(10).
				ShowLineNumbers(true).
				Value(&textInput).
				Run()
		}
		if optionsFile != "" {
			data, err := os.ReadFile(optionsFile)
			if err != nil {
				fmt.Printf("not able to read %s, the options will not be added to the db", optionsFile)
			} else {
				fileInput = string(data)
			}
		}

		optionLabels = getListInitOptions(textInput, fileInput)

		if len(optionLabels) > 0 {
			_, err = addOptions(db, optionLabels)
			if err != nil {
				fmt.Println("not able to add options to the database", err)
			} else {
				fmt.Printf("added %d options to the list\n", len(optionLabels))
			}
		}

	},
}

func init() {
	listCmd.AddCommand(listInitCmd)
	listInitCmd.Flags().BoolVarP(&optionsInput, "options", "o", false, "add options with a text input")
	listInitCmd.Flags().StringVarP(&optionsFile, "options-file", "f", "", "read options from a file")
}
