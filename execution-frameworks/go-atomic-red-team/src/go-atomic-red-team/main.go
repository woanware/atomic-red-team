package main

import (
	"fmt"
	"log"

	util "github.com/woanware/goutil"
	yaml "gopkg.in/yaml.v2"
)

func main() {

	atomic := new(Atomic)
	data, err := util.ReadTextFromFile("C:\\Dev\\atomic-red-team\\atomics\\T1002\\T1002.yaml")
	if err != nil {
		log.Fatalf("Error reading atomic: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &atomic)
	if err != nil {
		log.Fatalf("Error unmarshalling atomic: %v", err)
	}

	fmt.Println(atomic.AttackTechnique)
	fmt.Println(atomic.DisplayName)
}
