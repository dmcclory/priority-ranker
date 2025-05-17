/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"gorm.io/gorm"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var choice uint

func getPrompt(listData ListConfig) string {
	prompt := os.Getenv("RANKER_PROMPT")
	if prompt != "" {
		return prompt
	}
	globalPrompt := listData.GlobalPrompt
	if globalPrompt != "" {
		return globalPrompt
	}

	return "Which is more important?"
}

var loopVoting bool
var keepLooping bool

func vote(db *gorm.DB, options []Option, listData ListConfig) {
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	option1 := options[0]
	option2 := options[1]

	prompt := getPrompt(listData)

	if loopVoting {
		votes, err := loadVotes(db)
		if err != nil {
			fmt.Errorf("error while trying to get the vote count: %s", err)
			return
		}
		prompt = prompt + fmt.Sprintf(" (%d total votes)", len(votes))
	}

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
	err := form.Run()

	if err != nil {
		if loopVoting == true {
			fmt.Println("ok! - done looping!")
			keepLooping = false
		} else {
			fmt.Println("you cancelled the vote! - not voting!")
		}
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
	}
}

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Choose between two random options",
	Long: `
Two options will be randomly drawn from the active list.
You can choose which is more important (according to whatever prioritization criteria you like).
The choice will be recorded and used as part of the ranking computation

Use the '--loop' flag to vote repeatedly (Ctrl-C to stop)
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vote called and loopVoting is?", loopVoting)
		listData := loadLists()

		db, err := loadDb(dbPath(listData.ActiveList))
		check(err)

		options, err := loadOptions(db)
		check(err)

		if len(options) < 2 {
			fmt.Println("There aren't enough options to rank! Add more")
			return
		}

		if loopVoting {
			keepLooping = true
			for keepLooping == true {
				vote(db, options, listData)
			}
		} else {
			vote(db, options, listData)
			fmt.Println("Your vote has been recorded!")
		}
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)
	voteCmd.Flags().BoolVarP(&loopVoting, "loop", "l", false, "loop voting")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
