package cmd

import (
	"os"
	"strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func loadActiveProject() string {
	return "first-list"
}

func loadActiveProjectFromFile() string {
	data, err := os.ReadFile("/Users/dan/.ranker/active-list")
	check(err)

	return strings.TrimSpace(string(data))
}

func updateActiveList(listId string) {
	err := os.WriteFile("/Users/dan/.ranker/active-list", []byte(listId), 0644)
	check(err)
}
