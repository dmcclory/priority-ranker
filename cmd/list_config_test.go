package cmd

import (
	"errors"
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
