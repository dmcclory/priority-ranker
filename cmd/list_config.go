package cmd

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
	"strings"
)

type ChoiceList struct {
	Name string
	Id string
	Active bool
	DbExists bool
	CreatedAt int
}

func configPath() string {
	homedir, err := os.UserHomeDir()
	check(err)
	return path.Join(homedir, ".ranker", "config.json")
}

func dbPath(listId string) string {
	homedir, err := os.UserHomeDir()
	check(err)
	return path.Join(homedir, ".ranker", listId + ".sqlite")
}

type ListConfig struct {
	ActiveList string `json:"active"`
	Lists map[string]ChoiceList `json:"lists"`
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func updateActiveList(listId string, lists ListConfig) {
	lists.ActiveList = listId
	persistListConfig(lists)
}

func persistListConfig(lists ListConfig) {
	data, err := json.Marshal(lists)
	check(err)
	err = os.WriteFile(configPath(), data, 0644)
	check(err)
}

func markListEntryAsActive(listId string, lists ListConfig) {
	// optional: mark all others as false
	listCopy := lists.Lists[listId]
	listCopy.Active = true
	lists.Lists[listId] = listCopy
}

func markListEntryAsHavingDb(listId string, lists ListConfig) {
	var exists bool

	_, err := os.Stat(dbPath(listId))

	if os.IsNotExist(err) {
		exists = false
	} else {
		check(err)
		exists = true
	}

	listCopy := lists.Lists[listId]
	listCopy.DbExists = exists
	lists.Lists[listId] = listCopy
}

func loadLists() ListConfig {
	data, err := os.ReadFile(configPath())
	check(err)

	var result ListConfig

	// whoa -> this is important -> if you don't check that it fails
	// you'll get an empty object, that is no good!
	err = json.Unmarshal(data, &result)
	check(err)

	for k, _ := range result.Lists {
		markListEntryAsHavingDb(k, result)
	}

	return result
}

func generateId(listName string) string {

	// https://pkg.go.dev/regexp/syntax
	re := regexp.MustCompile(`[^[:alnum:]]+`)

	// https://pkg.go.dev/regexp#Regexp.ReplaceAllString
	id := strings.ToLower(re.ReplaceAllString(listName, "-"))

	return id
}
