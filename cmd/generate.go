package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Name string
	On   map[string]interface{}
	Jobs map[string]Job
}

type Job struct {
	RunsOn string
	Steps  []Step
}

type Step struct {
	Name string `yaml:"name"`
	Uses string `yaml:"uses,omitempty"`
	Run  string `yaml:"run,omitempty"`
}

var steps int

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a GitHub actions workflow",
	RunE: func(cmd *cobra.Command, args []string) error {
		var generatedSteps []Step
		generatedSteps = append(generatedSteps, Step{Name: "Checkout", Uses: "actions/checkout@v4"})
		for i := 1; i <= steps; i++ {
			generatedSteps = append(generatedSteps, Step{
				Name: fmt.Sprintf("Step %d", i),
				Run:  fmt.Sprintf("echo running step %d", i),
			})
		}

		wf := Workflow{
			Name: "Test",
			On: map[string]interface{}{
				"push": map[string][]string{
					"branches": {"main"},
				},
			},
			Jobs: map[string]Job{
				"build": {
					RunsOn: "ubuntu-latest",
					Steps:  generatedSteps,
				},
			},
		}

		out, err := yaml.Marshal(&wf)
		if err != nil {
			return err
		}
		err = os.WriteFile("workflow.yml", out, 0644)
		if err != nil {
			return err
		}

		fmt.Println("workflow.yml generated")
		return nil
	},
}

func init() {
	generateCmd.Flags().IntVarP(&steps, "steps", "s", 1, "Number of steps to generate")
	rootCmd.AddCommand(generateCmd)
}
