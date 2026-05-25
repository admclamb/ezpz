package cmd

import (
	"fmt"

	"github.com/admclamb/ezpz/internal/scaffold"
	"github.com/admclamb/ezpz/internal/stack"
	"github.com/spf13/cobra"
)

var (
	name    string
	stackID string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project from a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		reg := stack.NewDefaultRegistry()

		if name == "" {
			return fmt.Errorf("missing --name")
		}

		var chosen stack.Stack
		if stackID != "" {
			s, ok := reg.GetLatest(stackID)
			if !ok {
				return fmt.Errorf("unknown stack %q", stackID)
			}
			chosen = s
		} else {
			options := reg.ListLatest()
			if len(options) == 0 {
				return fmt.Errorf("no stacks available")
			}
			chosen = options[0]
		}

		return scaffoldProject(name, chosen)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&name, "name", "", "project name")
	initCmd.Flags().StringVar(&stackID, "stack", "", "template stack id")
}

func scaffoldProject(name string, chosen stack.Stack) error {
	return scaffold.Project(name, chosen)
}
