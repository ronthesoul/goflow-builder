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

	if _, err := os.Stat("workflow1.yml"); os.IsNotExist(err) {
		t.Fatalf("workflow not generated")
	}

	_ = os.Remove("workflow.yml")
}

func TestGenerateInvalidRun(t *testing.T) {
	rootCmd.SetArgs([]string{"generate", "--steps", "invalid"})
	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error for invalid --steps input, but got none")
	} else {
		t.Logf("caught expected error: %v", err)
	}

}
