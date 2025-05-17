package cmd

import (
	"errors"
	"os"
	"testing"
)

func getListsAsMap() ListConfig {
	return ListConfig{
		ActiveList: "first-list",
		Lists: map[string]OptionList{
			"first-list":  {Id: "first-list", Name: "First List"},
			"second-list": {Id: "second-list", Name: "Second List"},
			"third-list":  {Id: "third-list", Name: "Third List"},
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

	lists, err := addNewOptionList(lists, listName)

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

	_, err := addNewOptionList(lists, listName)

	if !errors.Is(err, OptionListExists) {
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

func TestCanAddANewListConfigToAnEmptyConfigStruct(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	lists := loadLists()

	if len(lists.Lists) != 0 {
		t.Errorf("Adding to empty state test, expected Lists to be an empty map")
	}

	lists, err = addNewOptionList(lists, "Cool New List")

	if len(lists.Lists) != 1 {
		t.Errorf("Adding to empty state test, expected Lists to have one item")
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

	markListEntryAsHavingDb("first-list", lists)
	markListEntryAsHavingDb("second-list", lists)

	if lists.Lists["first-list"].DbExists == false {
		t.Errorf("TestMarkEntryAsHavingDb - first-list was expected to be true but was false")
	}

	if lists.Lists["second-list"].DbExists == true {
		t.Errorf("TestMarkEntryAsHavingDb - second-list was expected to be false but was true")
	}
}

func TestCreateEmptyDbSavesAFileWithTheId(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	_, err = initDb(dbPath("test-db"))
	check(err)

	if fileDoesNotExist(dbPath("test-db")) {
		t.Errorf("expected to find a database for test-db in the temp directory, but none was found")
	}

	persistListConfig(getListsAsMap())
}

func TestDeleteListWithInvalidListId(t *testing.T) {
	lists := loadLists()

	_, err := deleteList(lists, "missing-id")

	if !errors.Is(err, OptionListMissing) {
		t.Errorf("Expected to get a OptionListMissing error, but did not")
	}
}

func TestDeleteListRemovesEntryAndFile(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)

	lists := getListsAsMap()

	_, err = initDb(dbPath("test-db"))
	check(err)

	lists, err = deleteList(lists, "first-list")

	_, secondListPresent := lists.Lists["second-list"]

	if !secondListPresent {
		t.Errorf("Expected the unaffected lists to remain in the set of lists in the config, but it was not")
	}


	if err != nil {
		t.Errorf("Expected err to be nil, but got %s", err)
	}

	if fileExists(dbPath("first-list")) {
		t.Errorf("Expected the list's database file to be deleted, but it was not")
	}

	_, firstListPresent := lists.Lists["first-list"]

	if firstListPresent {
		t.Errorf("Expected the list to be removed from the set of lists in the config, but it was not")
	}

	if lists.ActiveList != "" {
		t.Errorf("Expected the active list value to be cleared out, but it was not")
	}
}

func TestGetGlobalPromptDefaultsToBlank(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)
	persistListConfig(getListsAsMap())
	lists := loadLists()

	globalPrompt := getGlobalPrompt(lists)

	if globalPrompt != "" {
		t.Error("expected global prompt to eq: 'what do you think?'")
	}
}

func TestSettingGlobalPrompt(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "test")
	check(err)
	t.Setenv("RANKER_DIR", tempdir)
	persistListConfig(getListsAsMap())
	lists := loadLists()

	lists = setGlobalPrompt(lists, "What should the new prompt be?")

	globalPrompt := getGlobalPrompt(lists)

	if globalPrompt != "What should the new prompt be?" {
		t.Error("expected global prompt to eq: 'What should the new prompt be?', but got:", globalPrompt)
	}
}
