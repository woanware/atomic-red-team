package main

//
type Atomic struct {
	AttackTechnique string `yaml:"attack_technique"`
	DisplayName     string `yaml:"display_name"`
	AtomicTests     []struct {
		Name               string   `yaml:"name"`
		Description        string   `yaml:"description"`
		SupportedPlatforms []string `yaml:"supported_platforms"`
		InputArguments     struct {
			InputFile struct {
				Description string `yaml:"description"`
				Type        string `yaml:"type"`
				Default     string `yaml:"default"`
			} `yaml:"input_file"`
			OutputFile struct {
				Description string `yaml:"description"`
				Type        string `yaml:"type"`
				Default     string `yaml:"default"`
			} `yaml:"output_file"`
		} `yaml:"input_arguments,omitempty"`
		Executor struct {
			Name    string `yaml:"name"`
			Command string `yaml:"command"`
		} `yaml:"executor"`
	} `yaml:"atomic_tests"`
}
