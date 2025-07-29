package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the version of the application",
	Run: func(cmd *cobra.Command, args []string) {

		if Version == "" {
			fmt.Println("goflow-builder version: unknown")
		} else {
			fmt.Println("goflow-builder version:", Version)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
