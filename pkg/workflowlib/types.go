package workflowlib

type Workflow struct {
	Name string
	On   map[string]interface{}
	Jobs map[string]Job
}

type Job struct {
	RunsOn string
	Steps  []Step
}

type Step struct {
	Name string `yaml:"name"`
	Uses string `yaml:"uses,omitempty"`
	Run  string `yaml:"run,omitempty"`
}
