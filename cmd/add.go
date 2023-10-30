package cmd

import (
	"github.com/anurag-dhamala/go-task-cli/db"
	"github.com/spf13/cobra"
	"log"
)

func GetAddCmd() *cobra.Command {

	var addKey string
	var addValue string

	var add = &cobra.Command{
		Use:   "add",
		Short: "Add a new task (name and description)",
		Run: func(cmd *cobra.Command, args []string) {
			db.Insert(addKey, addValue)
		},
	}

	add.Flags().StringVarP(&addKey, "task-name", "n", "", "Add a new todo task name")
	add.Flags().StringVarP(&addValue, "task-desc", "d", "", "Add a new todo task desc")

	if err := add.MarkFlagRequired("task-name"); err != nil {
		log.Fatal(err)
	}

	if err := add.MarkFlagRequired("task-desc"); err != nil {
		log.Fatal(err)
	}

	return add
}
