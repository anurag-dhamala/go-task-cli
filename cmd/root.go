package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use:   "go-task-cli",
	Short: "go-task-cli to add todos",
}

func init() {
	rootCommand.AddCommand(GetVersionCmd())
	rootCommand.AddCommand(GetListCmd())
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal("cannot execute")
	}
}
