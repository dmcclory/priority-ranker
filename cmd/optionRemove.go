/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"errors"
	"gorm.io/gorm"
	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
)

// optionRemoveCmd represents the optionRemove command
var optionRemoveCmd = &cobra.Command{
	Use:   "remove",
	Aliases: []string{"delete"},
	Short: "Remove an option from the active list of options",
	Long: `Removes an option from the active list of options. Pass the numeric id for the option you wish to remove.
You can see the option ids by running 'ranker options'.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("\nYou must provide an ID for an option to remove one. Look the output of `ranker options` to see the ids.")
			return
		}
		optionId, err := strconv.ParseUint(args[0], 10, 32)
		listData := loadLists()
		db, err := loadDb(dbPath(listData.ActiveList))
		check(err)

		option, err := getOption(db, optionId)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("\nNo option found with id %d.\n", optionId)
			return
		} else {
			check(err)
		}

		var confirm bool
		huh.NewConfirm().
		  Title("This is irreversible and will delete your votes as well as the option. Are you sure?").
			Affirmative("Yup!").
			Negative("No ty").
			Value(&confirm).Run()

		if confirm {
			err = removeOption(db, optionId)

			if err != nil {
				fmt.Printf("\nError while trying to delete %d: %v\n", optionId, err)
			} else {
				fmt.Printf("\nSuccessfully deleted option '%s', with id: %d\n", option.Label, optionId)
			}
		} else {
			fmt.Println("no worries, exiting!")
		}
	},
}

func init() {
	optionCmd.AddCommand(optionRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optionRemoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optionRemoveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
