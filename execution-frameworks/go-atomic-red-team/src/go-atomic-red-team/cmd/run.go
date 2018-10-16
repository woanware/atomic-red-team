package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	util "github.com/woanware/goutil"
	yaml "gopkg.in/yaml.v2"
)

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

// Stub function to allow both the run and generate code to follow same code route
func run(cmd *cobra.Command, args []string) {

	runOrGenerate(cmd, args, false)
}

// Perform the run or generate action
func runOrGenerate(cmd *cobra.Command, args []string, generateOnly bool) {

	validate(args)

	atomic := new(Atomic)
	data, err := util.ReadTextFromFile(args[0])
	if err != nil {
		log.Fatalf("Error reading atomic: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &atomic)
	if err != nil {
		log.Fatalf("Error unmarshalling atomic: %v", err)
	}

	fmt.Printf("Execution OS: %s\n\n", runtime.GOOS)

	for _, test := range atomic.AtomicTests {

		fmt.Printf("[******** BEGIN TEST *******]\n\n")
		fmt.Println("Name: " + test.Name)
		fmt.Println("Description: " + strings.TrimSpace(test.Description))

		switch runtime.GOOS {
		case "windows":
			if util.DoesStringSliceContain(test.SupportedPlatforms, "windows") == false {
				fmt.Printf("Status: Cancelled (Test not supported under Windows)\n\n")
				fmt.Printf("[******** END TEST ********]\n\n")
				continue
			}

		case "linux":
			if util.DoesStringSliceContain(test.SupportedPlatforms, "linux") == false {
				fmt.Printf("Status: Cancelled (Test not supported under Linux)\n\n")
				fmt.Printf("[******** END TEST ********]\n\n")
				continue
			}

		}

		if strings.Contains(strings.ToLower(test.Executor.Name), "manual") == true {
			fmt.Printf("Status: Cancelled (Unable to run manual tests)\n\n")
			fmt.Printf("[******** END TEST ********]\n\n")
			continue
		}

		finalCommand := test.Executor.Command

		// Transpose the command arguments
		if len(test.InputArguments) > 0 {
			for k := range test.InputArguments {

				i := test.InputArguments[k]
				finalCommand = strings.Replace(finalCommand, "#{"+k+"}", i.Default, 1)
			}
		}

		if generateOnly == true {
			fmt.Println("Command: " + finalCommand)
		}

		fmt.Println("Technique: " + atomic.AttackTechnique)
		fmt.Println("Display Name: " + atomic.DisplayName)

		fmt.Printf("\nDo you want to execute the test:")
		ret, err := util.GetYesNoPrompt(false)
		if err != nil {
			fmt.Printf("\nStatus: Error (%v)\n\n", err)
			fmt.Printf("[******** END TEST ********]\n\n")
			continue
		}

		if ret == false {
			fmt.Printf("\nStatus: Cancelled\n\n")
			fmt.Printf("[******** END TEST ********]\n\n")
			continue
		}

		fmt.Println("")
		processExecutor(finalCommand, test)
		fmt.Printf("[******** END TEST ********]\n\n")
	}
}

// Performs specific executor actions
func processExecutor(finalCommand string, test Test) {

	var err error
	commands := strings.Split(finalCommand, "\n")
	for _, command := range commands {

		switch test.Executor.Name {
		case "command_prompt":
			err = executeCmdCommand(command)

		case "powershell":
			err = executePowershellCommand(finalCommand)

		case "sh":
			err = executeShCommand(finalCommand)

		case "bash":
			err = executeBashCommand(finalCommand)

		default:
			fmt.Printf("Status: Cancelled (Unsupported executor)\n\n")
		}

		if err != nil {
			fmt.Printf("Status: Error (%v)\n\n", err)
		} else {
			fmt.Printf("Status: Success\n\n")
		}

	}
}

// Executes the command using cmd.exe
func executeCmdCommand(command string) error {

	fmt.Printf("Command: cmd /C %s\n", command)
	err := exec.Command("cmd", "/C", command).Start()
	if err != nil {
		return err
	}

	return nil
}

// Executes the command using Powershell
func executePowershellCommand(command string) error {

	fmt.Printf("Command: Invoke-Command -ScriptBlock %s\n", command)
	err := exec.Command("Invoke-Command", fmt.Sprintf("-ScriptBlock", command)).Start()
	if err != nil {
		return err
	}

	return nil
}

// Executes the command using Powershell
func executeShCommand(command string) error {

	fmt.Printf("Command:sh -c %s\n", command)
	err := exec.Command("sh", "-c", command).Start()
	if err != nil {
		return err
	}

	return nil
}

// Executes the command using Powershell
func executeBashCommand(command string) error {

	fmt.Printf("Command: bash -c %s\n", command)
	err := exec.Command("bash", "-c", command).Start()
	if err != nil {
		return err
	}

	return nil
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
