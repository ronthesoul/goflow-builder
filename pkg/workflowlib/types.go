package workflowlib

type Workflow struct {
	Name string
	On   Trigger
	Jobs map[string]Job
}

type Trigger struct {
	Push Push `yaml:"push"`
}

type Push struct {
	Branches []string `yaml:"branches"`
}

type Job struct {
	RunsOn string            `yaml:"runs-on,omitempty"`
	Env    map[string]string `yaml:"env,omitempty"`
	Steps  []Step            `yaml:"steps,omitempty"`
}

type Step struct {
	Name string `yaml:"name,omitempty"`
	Uses string `yaml:"uses,omitempty"`
	Run  string `yaml:"run,omitempty"`
}
