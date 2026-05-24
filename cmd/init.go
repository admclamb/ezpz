package cmd

import (
	"fmt"
	"os"

	"github.com/admclamb/ezpz/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	use: "init",
	Short: "Initialize the registry",
	Run: func(cmd *cobra.Command, args []string) {
		
	}
}