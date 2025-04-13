package cmd

import (
	"encoding/json"
	"os"
)

type ChoiceList struct {
	Name string
	Id string
	Active bool
	DbExists bool
}

type ListConfig struct {
	ActiveList string `json:"active"`
	Lists []ChoiceList `json:"lists"`
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func updateActiveList(listId string, lists ListConfig) {
	lists.ActiveList = listId
	data, err := json.Marshal(lists)
	check(err)
	err = os.WriteFile("/Users/dan/.ranker/config.json", data, 0644)
	check(err)
}

func loadLists() ListConfig {
	data, err := os.ReadFile("/Users/dan/.ranker/config.json")
	check(err)

	var result ListConfig

	// whoa -> this is important -> if you don't check that it fails
	// you'll get an empty object, that is no good!
	err = json.Unmarshal(data, &result)
	check(err)

	return result
}
