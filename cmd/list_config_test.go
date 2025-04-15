package cmd

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func getListsAsMap() ListConfig{
	return ListConfig{
		ActiveList: "first-list",
		Lists: map[string]ChoiceList{
			"first-list": {Id: "first-list", Name: "First List"},
			"second-list": {Id: "second-list", Name: "Second List"},
			"third-list": {Id: "third-list", Name: "Third List"},
		},
	}
}

func TestMarkActiveMap(t *testing.T) {
	lists := getListsAsMap()

	markListEntryAsActive(lists.Lists["second-list"].Id, lists)

	if lists.Lists["second-list"].Active != true {
		t.Errorf("markActiveListEntry failed to update the correct list")
	}
}

func TestGenerateId(t *testing.T) {
	id := generateId("Input Is Good")

	if id != "input-is-good" {
		t.Errorf("TestGenerateId failed. expected: 'input-is-good', actual: %s", id)
	}
}

func TestAddNewChoiceHappyPath(t *testing.T) {
	lists := getListsAsMap()
	listName := "Fourth List"

	lists, err := addNewChoiceList(lists, listName)

	check(err)

	_, newOptionPresent := lists.Lists["fourth-list"]

	if newOptionPresent != true {
		t.Errorf("TestAddNewChoiceHappyPath failed to add the new list to the set of lists")
	}

	if lists.ActiveList != "fourth-list" {
		t.Errorf("TestAddNewChoiceHappyPath failed to set the new list to the active list")
	}
}

func TestAddNewChoiceProjectAlreadyExists(t *testing.T) {
	lists := getListsAsMap()
	listName := "First List"

	_, err := addNewChoiceList(lists, listName)

	if !errors.Is(err, ChoiceListExists) {
		t.Errorf("Error expected, but the function returned success")
	}

}

func TestSettingHomeDir(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	persistListConfig(getListsAsMap())

	lists := loadLists()

	if lists.ActiveList != "first-list" {
		t.Errorf("SettingHomeDir - does not have expected active list ID")
	}
}

func TestLoadListsWhenConfigDoesNotExistYet(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)
	
	lists := loadLists()

	if lists.ActiveList != "" {
		t.Errorf("Loading Empty state test, expected ActiveList to be blank")
	}

	if len(lists.Lists) != 0 {
		t.Errorf("Loading Empty state test, expected Lists to be an empty map")
	}
}

func TestMarkEntryAsHavingDb(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	persistListConfig(getListsAsMap())

	err = os.WriteFile(dbPath("first-list"), []byte("data"), 0644)
	check(err)

	lists := loadLists()

	fmt.Println("got this", lists)

	markListEntryAsHavingDb("first-list", lists)
	markListEntryAsHavingDb("second-list", lists)

	if lists.Lists["first-list"].DbExists == false {
		t.Errorf("TestMarkEntryAsHavingDb - first-list was expected to be true but was false")
	}

	if lists.Lists["second-list"].DbExists == true {
		t.Errorf("TestMarkEntryAsHavingDb - second-list was expected to be false but was true")
	}
}
