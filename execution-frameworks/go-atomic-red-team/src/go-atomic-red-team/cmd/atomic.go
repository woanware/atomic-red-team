package cmd

// Encapsulates an input argument from the atomic file
type Input struct {
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Default     string `yaml:"default"`
}

type Test struct {
	Name               string           `yaml:"name"`
	Description        string           `yaml:"description"`
	SupportedPlatforms []string         `yaml:"supported_platforms"`
	InputArguments     map[string]Input `yaml:"input_arguments,omitempty"`
	//InputArguments map[string]interface{} `yaml:"input_arguments,omitempty"`
	Executor struct {
		Name    string `yaml:"name"`
		Command string `yaml:"command"`
	} `yaml:"executor"`
}

// Encapsulates an atomic
type Atomic struct {
	AttackTechnique string `yaml:"attack_technique"`
	DisplayName     string `yaml:"display_name"`
	AtomicTests     []Test `yaml:"atomic_tests"`
}
