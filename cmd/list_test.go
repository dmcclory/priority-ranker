package cmd

import (
	"strings"
	"testing"
)

func TestGetLists(t *testing.T) {
	res := getLists()

	if res[0].Id != "first-list" {
		t.Errorf("no! fail!")
	}
}

func TestListNames(t *testing.T) {
	lists := getLists()

	names := getListNames(lists)

	if names[0] != "First List" {
		t.Errorf("the test failed")
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
