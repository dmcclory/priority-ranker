package cmd

import (
	"testing"
)

func TestGetListInitOptionsReturnsAnArrayOfStrings(t *testing.T) {

	huhInput := `Input 1
Input 2
Input 3`
  fileInput := `File 1
File 2
File 3`

  options := getListInitOptions(huhInput, fileInput)

	if len(options) != 6 {
		t.Errorf("Expected getListInitOptions to return 6 entries, but got %d", len(options))
	}
}

func TestGetListInitOptionsFiltersOutBlankLines(t *testing.T) {

	huhInput := `Input 1

Input 2
Input 3
`
  fileInput := `File 1
File 2
File 3
`

  options := getListInitOptions(huhInput, fileInput)

	if len(options) != 6 {
		t.Errorf("Expected getListInitOptions to return 6 entries, but got %d", len(options))
	}
}

func TestGetListInitReturnsSingleInstancesOfDuplicateLines(t *testing.T) {

	huhInput := `Input 1
Input 2
Input 3
Input 3
Input 3
`
  fileInput := `File 1
File 2
File 3
`

  options := getListInitOptions(huhInput, fileInput)

	if len(options) != 6 {
		t.Errorf("Expected getListInitOptions to return 6 entries, but got %d", len(options))
	}
}
