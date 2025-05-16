package cmd

import (
	// "errors"
	// "os"
	"testing"
)

func TestGetPromptWithNoEnvVarOrSettingUsesDefault(t *testing.T) {
	res := getPrompt()

	if res != "Which is more important?" {
		t.Errorf("Expected the default prompt to be 'Which is more important?'")
	}
}

func TestGetPromptWithEnvVarOverride(t *testing.T) {
	t.Setenv("RANKER_PROMPT", "Custom prompt")
	res := getPrompt()

	if res != "Custom prompt" {
		t.Errorf("Expected the default prompt to be 'Custom prompt'")
	}
}
