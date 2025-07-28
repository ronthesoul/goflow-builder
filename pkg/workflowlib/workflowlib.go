package workflowlib

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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

func CreateWorkflow(name string, steps []Step) Workflow {
	return Workflow{
		Name: name,
		On: map[string]interface{}{
			"push": map[string][]string{
				"branches": {"main"},
			},
		},
		Jobs: map[string]Job{
			"build": {
				RunsOn: "ubuntu-latest",
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
