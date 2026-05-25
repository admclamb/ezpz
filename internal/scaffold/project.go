package scaffold

import (
	"fmt"

	"github.com/admclamb/ezpz/internal/stack"
)

func Project(name string, chosen stack.Stack) error {
	if name == "" {
		return fmt.Errorf("project name is required")
	}

	if chosen.ID == "" {
		return fmt.Errorf("stack id is required")
	}

	return fmt.Errorf("TODO: scaffold project %q from template %q (%s)", name, chosen.ID, chosen.TemplateRepo)
}
