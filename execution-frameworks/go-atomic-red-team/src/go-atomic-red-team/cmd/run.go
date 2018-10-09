package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	util "github.com/woanware/goutil"
)

// ##### Constants ###########################################################

const APP_TITLE string = "go-atomic-red-team"
const APP_NAME string = "gart"
const APP_VERSION string = "v0.0.1"

// ##### Variables ###########################################################

var cmdRun = &cobra.Command{
	Use:     "run [atomic path]",
	Aliases: []string{"atomic"},
	Short:   "Runs an atomic",
	Long:    `Runs an atomic using the yaml file specified`,
	Run:     run,
}

// ##### Functions ###########################################################

// Add the command to the cobra setup
func init() {

	cmdRoot.AddCommand(cmdRun)
}

// Perform the run action
func run(cmd *cobra.Command, args []string) {

	validate()

	atomic := new(Atomic)
	data, err := util.ReadTextFromFile(args[0])
	if err != nil {
		log.Fatalf("Error reading atomic: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &atomic)
	if err != nil {
		log.Fatalf("Error unmarshalling atomic: %v", err)
	}

	fmt.Println("Running under " + runtime.GOOS)

	for _, test := range atomic.AtomicTests {

		fmt.Println(test.Name)
		fmt.Println(test.Description)

		switch runtime.GOOS {
		case "windows":
			if util.DoesStringSliceContain(test.SupportedPlatforms, "windows") == false {
				fmt.Println("Test not supported under Windows")
				continue
			}

		case "linux":
			if util.DoesStringSliceContain(test.SupportedPlatforms, "linux") == false {
				fmt.Println("Test not supported under Linux")
				continue
			}

		}

		if strings.Contains(strings.ToLower(test.Executor.Name), "manual") == true {
			fmt.Println("Unable to run manual tests")
			continue
		}

		finalCommand := test.Executor.Command

		if len(test.InputArguments) > 0 {
			fmt.Println("Replacing input arguments with default values")

			for k, v := range test.InputArguments {

				i := test.InputArguments[k]
				finalCommand = strings.Replace(finalCommand, "#{"+k+"}", i.Default, 1)
			}
		}

		fmt.Println("[********BEGIN TEST*******]")
		fmt.Println(atomic.AttackTechnique)
		fmt.Println(atomic.DisplayName)
	}
}

// Validate the parameters passed to the action
func validate(args []string) {

	if len(args) == 0 {
		fmt.Println("Path to atomic file not supplied")
		os.Exit(0)
	} else if len(args) > 1 {
		fmt.Println("Only the path to atomic file should be supplied")
		os.Exit(0)
	}

	if util.DoesFileExist(args[0]) == false {
		fmt.Println("Atomic file does not exist")
		os.Exit(1)
	}
}
