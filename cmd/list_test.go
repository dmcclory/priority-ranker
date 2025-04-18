package cmd

import (
	"strings"
	"testing"
)

func getLists() []OptionList {
	return []OptionList{
		{Id: "first-list", Name: "First List"},
		{Id: "second-list", Name: "Second List"},
		{Id: "third-list", Name: "Third List"},
	}
}

func getEmptyListResult() []OptionList {
	return []OptionList{}
}

func TestGetLists(t *testing.T) {
	res := getLists()

	if res[0].Id != "first-list" {
		t.Errorf("no! fail!")
	}
}

func TestEmptyStateMessage(t *testing.T) {
	msg := emptyStateMessage()

	if strings.HasPrefix(msg, "no lists have been created") == false {
		t.Errorf("empty state message text is incorrect")
	}
}
