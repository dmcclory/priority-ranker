/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var choice uint

func getPrompt() string {
	prompt := os.Getenv("RANKER_PROMPT")
	if prompt != "" {
		return prompt
	} else {
		return "Which is more important?"
	}
}

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Choose between two random options",
	Long: `
Two options will be randomly drawn from the active list.
You can choose which is more important (according to whatever prioritization criteria you like).
The choice will be recorded and used as part of the ranking computation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vote called")
		listData := loadLists()

		db, err := loadDb(dbPath(listData.ActiveList))
		check(err)

		options, err := loadOptions(db)
		check(err)

		if len(options) < 2 {
			fmt.Println("There aren't enough options to rank! Add more")
			return
		}

		rand.Shuffle(len(options), func(i, j int) {
			options[i], options[j] = options[j], options[i]
		})

		option1 := options[0]
		option2 := options[1]

		prompt := getPrompt()

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[uint]().
					Title(prompt).
					Options(
						huh.NewOption(option1.Label, option1.ID),
						huh.NewOption(option2.Label, option2.ID),
					).
				Value(&choice),
			),
		)

		err = form.Run()

		if err != nil {
			fmt.Println("you cancelled the vote! - not voting!")
			return
		}

		var winnerId, loserId uint
		if choice == option1.ID {
			winnerId = option1.ID
			loserId = option2.ID
		} else {
			winnerId = option2.ID
			loserId = option1.ID
		}
		err = addVote(db, winnerId, loserId)
		if err != nil {
			fmt.Printf("We were not able to save your vote because: %e\n", err)
		} else {
			fmt.Println("Your vote has been recorded!")
		}
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
