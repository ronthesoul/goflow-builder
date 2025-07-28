package cmd

import (
	"fmt"

	"github.com/ronthesoul/workflow-builder/pkg/workflowlib"
	"github.com/spf13/cobra"
)

var steps int
var file string
var name string
var notify bool
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
		wf := workflowlib.CreateWorkflow(name, env, generatedSteps)
		if err := workflowlib.WriteYAMLToFile(wf, file); err != nil {
			return err
		}
		fmt.Println("workflow.yml generated")
		return nil
	},
}

func init() {
	generateCmd.Flags().IntVarP(&steps, "steps", "s", 1, "Number of steps to generate")
	generateCmd.Flags().StringVarP(&file, "file", "f", "workflow.yml", "Output file")
	generateCmd.Flags().StringVarP(&name, "name", "n", "ga-workflow", "Name of the workflow")
	generateCmd.Flags().BoolVarP(&notify, "notify", "y", false, "Enable Slack notifications")
	rootCmd.AddCommand(generateCmd)
}
