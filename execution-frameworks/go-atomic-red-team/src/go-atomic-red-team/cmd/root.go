package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ##### Constants ###########################################################

const APP_TITLE string = "go-atomic-red-team"
const APP_NAME string = "gart"
const APP_VERSION string = "v0.0.1"

// ##### Variables ###########################################################

var cmdRoot = &cobra.Command{
	Use:   "gart",
	Short: "gart is an atomic red team execution engine",
	Long:  `gart executes atomic yaml files`,
	Run:   root,
}

// ##### Functions ###########################################################

// Root/base/default command e.g. just display the app info
func Execute() {

	fmt.Printf("\n%s (%s) %s\n\n", APP_TITLE, APP_NAME, APP_VERSION)

	if err := cmdRoot.Execute(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

// Core functionality for the root/base/default command
func root(cmd *cobra.Command, args []string) {

	// If no verb is supplied then just display the "help"
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
}
