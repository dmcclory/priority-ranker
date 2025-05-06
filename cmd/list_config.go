package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

type OptionList struct {
	Name      string
	Id        string
	Active    bool
	DbExists  bool
	CreatedAt int64
}

var OptionListExists = fmt.Errorf("Entry already exists in List config")
var OptionListMissing = fmt.Errorf("Entry does not exist in List config")

func rankerDir() string {
	// home
	rankerDir := os.Getenv("RANKER_DIR")

	if len(rankerDir) > 0 {
		return rankerDir
	}

	homedir, err := os.UserHomeDir()
	check(err)
	return path.Join(homedir, ".ranker")
}

func configPath() string {
	dir := rankerDir()
	return path.Join(dir, "config.json")
}

func dbPath(listId string) string {
	dir := rankerDir()
	return path.Join(dir, listId+".sqlite")
}

type ListConfig struct {
	ActiveList string                `json:"active"`
	Lists      map[string]OptionList `json:"lists"`
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

func fileDoesNotExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
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
	if fileDoesNotExist(configPath()) {
		return ListConfig{Lists: map[string]OptionList{}}
	}

	data, err := os.ReadFile(configPath())
	check(err)

	var result ListConfig

	// whoa -> this is important -> if you don't check that it fails
	// you'll get an empty object, that is no good!
	err = json.Unmarshal(data, &result)
	check(err)

	// for k, _ := range result.Lists {
	for k := range result.Lists {
		markListEntryAsHavingDb(k, result)
	}

	return result
}

func addNewOptionList(lists ListConfig, listName string) (ListConfig, error) {
	listId := generateId(listName)
	createdAt := time.Now().Unix()

	_, present := lists.Lists[listId]

	if present {
		return ListConfig{}, OptionListExists
	}

	newOptionList := OptionList{Name: listName, Id: listId, CreatedAt: createdAt}
	lists.Lists[listId] = newOptionList
	lists.ActiveList = listId

	return lists, nil
}

func createEmptyDb(listId string) error {
	path := dbPath(listId)

	err := os.WriteFile(path, []byte("data"), 0644)

	return err
}

func generateId(listName string) string {

	// https://pkg.go.dev/regexp/syntax
	re := regexp.MustCompile(`[^[:alnum:]]+`)

	// https://pkg.go.dev/regexp#Regexp.ReplaceAllString
	id := strings.ToLower(re.ReplaceAllString(listName, "-"))

	return id
}

func deleteList(lists ListConfig, listId string) (ListConfig, error) {
	_, listPresent := lists.Lists[listId]
	if !listPresent {
		return ListConfig{}, OptionListMissing
	}

	err := os.Remove(dbPath(listId))
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("not able to delete db for: ", listId)
	}

	if lists.ActiveList == listId {
		lists.ActiveList = ""
	}

	delete(lists.Lists, listId)

	return lists, nil
}
