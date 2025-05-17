package cmd

import (
	// "errors"
	// "os"
	"testing"
)

func TestGetPromptWithNoEnvVarOrSettingUsesDefault(t *testing.T) {
	listConfig := getListsAsMap()
	res := getPrompt(listConfig)

	if res != "Which is more important?" {
		t.Errorf("Expected the default prompt to be 'Which is more important?'")
	}
}

func TestGetPromptWithEnvVarOverride(t *testing.T) {
	listConfig := getListsAsMap()
	t.Setenv("RANKER_PROMPT", "Custom prompt")
	res := getPrompt(listConfig)

	if res != "Custom prompt" {
		t.Errorf("Expected the default prompt to be 'Custom prompt'")
	}
}

func TestGetPromptWithEnvVarGlobalPromptSet(t *testing.T) {
	listConfig := getListsAsMap()
	listConfig.GlobalPrompt = "A Global Prompt"
	res := getPrompt(listConfig)

	if res != "A Global Prompt" {
		t.Errorf("Expected the default prompt to be 'A Global Prompt'")
	}
}
