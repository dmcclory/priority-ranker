/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "math/rand"
	// "os"
	// "github.com/charmbracelet/huh"
	"strings"
	"github.com/spf13/cobra"
)

var configPromptSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Update the global prompt",
	Run: func(cmd *cobra.Command, args[]string) {
		lists := loadLists()
		newPrompt := strings.Join(args, " ")
		if newPrompt == "" {
			fmt.Printf("")
			return
		}
		lists = setGlobalPrompt(lists, newPrompt)
		fmt.Printf("\nnew global prompt:'%s'\n", lists.GlobalPrompt)
	},
}

func init() {
	configPromptCmd.AddCommand(configPromptSetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
