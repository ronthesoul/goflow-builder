package workflowlib

type Workflow struct {
	Name string
	On   map[string]interface{}
	Jobs map[string]Job
}

type Job struct {
	RunsOn string            `yaml:"runsOn,omitempty"`
	Steps  []Step            `yaml:"steps,omitempty"`
	Env    map[string]string `yaml:"env,omitempty"`
}

type Step struct {
	Name string `yaml:"name,omitempty"`
	Uses string `yaml:"uses,omitempty"`
	Run  string `yaml:"run,omitempty"`
}
