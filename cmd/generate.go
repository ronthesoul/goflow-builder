package cmd

import (
	"fmt"

	"github.com/ronthesoul/goflow-builder/pkg/workflowlib"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var steps int
var file string
var name string
var runner string
var notify bool
var dryRun bool
var env = map[string]string{}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a GitHub actions workflow",
	RunE: func(cmd *cobra.Command, args []string) error {
		if notify {
			env = workflowlib.AddEnvToYaml()
			if err := workflowlib.CopyNotifyFile(workflowlib.FindFolder(file)); err != nil {
				return err
			}
		}
		generatedSteps := workflowlib.BuildSteps(steps)
		wf := workflowlib.CreateWorkflow(name, runner, env, generatedSteps)
		if dryRun {
			yamlData, err := yaml.Marshal(wf)
			if err != nil {
				return err
			}
			fmt.Println(string(yamlData))
		} else {
			if err := workflowlib.WriteYAMLToFile(wf, file); err != nil {
				return err
			}
			fmt.Println("workflow.yml generated")
		}
		return nil
	},
}

func init() {
	generateCmd.Flags().IntVarP(&steps, "steps", "s", 1, "Number of steps to generate")
	generateCmd.Flags().StringVarP(&file, "file", "f", "workflow.yml", "Output file")
	generateCmd.Flags().StringVarP(&name, "name", "n", "ga-workflow", "Name of the workflow (Default: ga-workflow)")
	generateCmd.Flags().StringVarP(&runner, "runner", "r", "ubuntu-latest", "Choose runner OS (Default: ubuntu-latest)")
	generateCmd.Flags().BoolVarP(&notify, "notify", "y", false, "Enable Slack notifications")
	generateCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Output YAML to console instead of writing to file")
	rootCmd.AddCommand(generateCmd)
}
