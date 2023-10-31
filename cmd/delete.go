package cmd

import (
	"log"

	"github.com/anurag-dhamala/go-task-cli/db"
	"github.com/spf13/cobra"
)

func DeleteTaskByName() *cobra.Command {

	var taskName string
	var deleteTask = &cobra.Command{
		Use:   "delete",
		Short: "delete the completed task",
		Run: func(cmd *cobra.Command, args []string) {
			db.DeleteByKey(taskName)
		},
	}

	deleteTask.Flags().StringVarP(&taskName, "task-delete", "n", "", "Delete the task if you have completed")

	if err := deleteTask.MarkFlagRequired("task-delete"); err != nil {
		log.Fatal(err)
	}
	return deleteTask
}
