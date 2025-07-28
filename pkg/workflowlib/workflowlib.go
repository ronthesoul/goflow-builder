package workflowlib

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Generate functions

func BuildSteps(count int) []Step {
	steps := []Step{{Name: "Checkout", Uses: "actions/checkout@v4"}}
	for i := 1; i <= count; i++ {
		steps = append(steps, Step{
			Name: fmt.Sprintf("Step %d", i),
			Run:  fmt.Sprintf("echo running step %d", i),
		})
	}
	return steps
}

func CreateWorkflow(name string, env map[string]string, steps []Step) Workflow {
	return Workflow{
		Name: name,
		On: Trigger{
			Push: Push{
				Branches: []string{"main"},
			},
		},
		Jobs: map[string]Job{
			"build": {
				RunsOn: "ubuntu-latest",
				Env:    env,
				Steps:  steps,
			},
		},
	}
}

func WriteYAMLToFile(data interface{}, path string) error {
	out, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0644)
}

// Notify flag functions

func AddEnvToYaml() map[string]string {
	return map[string]string{
		"SLACK_WEBHOOK": "${{ secrets.SLACK_WEBHOOK }}",
		"NOTIFY":        ".github/workflows/notify.sh",
	}
}

func FindFolder(file string) string {
	absPath, _ := filepath.Abs(file)
	targetDir := filepath.Dir(absPath)
	return targetDir
}

func CopyNotifyFile(targetDir string) error {
	data, err := os.ReadFile("templates/notify/notify.sh")
	if err != nil {
		return err
	}
	dest := filepath.Join(targetDir, "notify.sh")
	return os.WriteFile(dest, data, 0644)
}
