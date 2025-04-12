package cmd

import (
	"strings"
	"testing"
)

func getLists() []ChoiceList {
	return []ChoiceList{
		{Id: "first-list", Name: "First List"},
		{Id: "second-list", Name: "Second List"},
		{Id: "third-list", Name: "Third List"},
	}
}

func getEmptyListResult() []ChoiceList {
	return []ChoiceList{}
}

func TestGetLists(t *testing.T) {
	res := getLists()

	if res[0].Id != "first-list" {
		t.Errorf("no! fail!")
	}
}

func TestMarkListNames(t *testing.T) {
	lists := getLists()

	lists[1].Active = false

	lists = markActiveListEntry(lists, lists[1].Id)

	if lists[1].Active != true {
		t.Errorf("markActiveListEntry failed to update the correct list")
	}
}

func TestEmptyStateMessage(t *testing.T) {
	msg := emptyStateMessage()

	if strings.HasPrefix(msg, "no lists have been created") == false {
		t.Errorf("empty state message text is incorrect")
	}
}
