/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

type RankedResult struct {
	Rank uint
	Label string
	Score float64
}

func getRankedResults() ([]RankedResult, error) {
	listData := loadLists()
	db, err := loadDb(dbPath(listData.ActiveList))
	check(err)
	options, err := loadOptions(db)
	check(err)
	votes, err := loadVotes(db)
	check(err)

	if err != nil {
		return []RankedResult{}, err
	}

	optionIds := []uint{}
	pScores := PScores{}

	for _, option := range options {
		optionIds = append(optionIds, option.ID)
		pScores[option.ID] = 1
	}

  winRecord := buildWinRecordFromVotes(votes, optionIds)

	ranks := calcNewPScores(winRecord, pScores)

	sort.Slice(options, func(i, j int) bool {
		if ranks[options[i].ID] > ranks[options[j].ID] {
			return true
		} else {
			return false
		}
	})

	rankedResults := []RankedResult{}

	for i, option := range options {
		rankedResults = append(rankedResults, RankedResult{
	    Rank: uint(i),
			Label: option.Label,
			Score: ranks[option.ID],
		})
	}

	return rankedResults, nil
}

var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "See the ranking of the options",
	Long: `Prints a table of the relative ranking of the options, based on your vote history.

	It takes ~30 votes to start to get meaningful data.`,
	Aliases: []string{"result", "rankings"},
	Run: func(cmd *cobra.Command, args []string) {
		defaultStyles := table.DefaultStyles()

		styles := table.Styles{
			Cell:   defaultStyles.Cell,
			Header: defaultStyles.Header,
		}

		results, err := getRankedResults()
		check(err)
		var rows [][]string

		for _, result := range results {
			rows = append(rows, []string{fmt.Sprintf("%d", result.Rank + 1), result.Label, fmt.Sprintf("%f", result.Score)})
		}

		table := ltable.New().Headers("Rank", "Option", "Score").
		  Rows(rows...).
			StyleFunc(func(row, _ int) lipgloss.Style {
				if row == 0 {
					return styles.Header
				}
				return styles.Cell
			})
		fmt.Println(table.Render())
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resultsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resultsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
