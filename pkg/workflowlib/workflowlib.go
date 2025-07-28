package workflowlib

import (
	"fmt"
)

func GenerateStep(name string, index int) string {
	return fmt.Sprintf("Step %d - %s", index, name)
}
