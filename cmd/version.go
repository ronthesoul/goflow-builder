/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
		fmt.Println("goflow-builder version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
