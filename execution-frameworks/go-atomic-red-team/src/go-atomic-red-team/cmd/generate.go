package cmd

import (
	"github.com/spf13/cobra"
)

// ##### Variables ###########################################################

var cmdGenerate = &cobra.Command{
	Use:     "generate [atomic path]",
	Aliases: []string{"atomic"},
	Short:   "Generates an atomic",
	Long:    `Generates an atomic using the yaml file specified e.g. only shows commands`,
	Run:     run,
}

// ##### Functions ###########################################################

// Add the command to the cobra setup
func init() {

	cmdRoot.AddCommand(cmdGenerate)
}

// Stub function to allow both the run and generate code to follow same code route
func generate(cmd *cobra.Command, args []string) {

	runOrGenerate(cmd, args, true)
}
