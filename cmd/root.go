package cmd

import (
	"fmt"
	"os"

	"github.com/admclamb/ezpz/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ezpz",
	Short: "ezpz is an project automation tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.Greet("World"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
