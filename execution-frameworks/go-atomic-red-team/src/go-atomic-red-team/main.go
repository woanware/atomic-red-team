package main

import (
	"go-atomic-red-team/cmd"
)

// ##### Functions ###########################################################

// Application entry point
func main() {
	cmd.Execute()
}

// func main() {

// 	fmt.Printf("\n%s (%s) %s\n\n", APP_TITLE, APP_NAME, APP_VERSION)

// 	atomic := new(Atomic)
// 	data, err := util.ReadTextFromFile("C:\\Dev\\atomic-red-team\\atomics\\T1002\\T1002.yaml")
// 	if err != nil {
// 		log.Fatalf("Error reading atomic: %v", err)
// 	}

// 	err = yaml.Unmarshal([]byte(data), &atomic)
// 	if err != nil {
// 		log.Fatalf("Error unmarshalling atomic: %v", err)
// 	}

// 	fmt.Println("Running under " + runtime.GOOS)

// 	for _, test := range atomic.AtomicTests {

// 		fmt.Println(test.Name)
// 		fmt.Println(test.Description)

// 		fmt.Println(len(test.InputArguments))

// 		switch runtime.GOOS {
// 		case "windows":
// 			if util.DoesStringSliceContain(test.SupportedPlatforms, "windows") == false {
// 				fmt.Println("Test not supported under Windows")
// 				continue
// 			}

// 		case "linux":
// 			if util.DoesStringSliceContain(test.SupportedPlatforms, "linux") == false {
// 				fmt.Println("Test not supported under Linux")
// 				continue
// 			}

// 		}

// 		if strings.Contains(strings.ToLower(test.Executor.Name), "manual") == true {

// 		}

// 		finalCommand := test.Executor.Command

// 		fmt.Println(len(test.InputArguments))

// 		if len(test.InputArguments) > 0 {
// 			fmt.Println("Replacing input arguments with default values")

// 			for k, v := range test.InputArguments {

// 				i := test.InputArguments[k]

// 				//fmt.Println(i.Default)
// 				//fmt.Println(i.Description)
// 				//fmt.Println(i.Type)

// 				finalCommand = strings.Replace(finalCommand, "#{"+k+"}", i.Default, 1)
// 			}
// 		}

// 		//Write-Debug -Message 'Getting executor and build command script'

// 		fmt.Println("[********BEGIN TEST*******]")
// 		fmt.Println(atomic.AttackTechnique)
// 		fmt.Println(atomic.DisplayName)
// 	}

// }
