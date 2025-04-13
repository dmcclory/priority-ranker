package cmd

import (
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
