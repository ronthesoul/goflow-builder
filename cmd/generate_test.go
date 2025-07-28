package cmd

import (
	"os"
	"testing"
)

func TestGenerateWorkflowFile(t *testing.T) {
	steps = 2

	cmd := generateCmd
	err := cmd.RunE(cmd, []string{})

	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	if _, err := os.Stat("workflow.yml"); os.IsNotExist(err) {
		t.Fatalf("workflow not generated")
	}

	_ = os.Remove("workflow.yml")

}
