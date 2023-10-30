package cmd

import (
	"github.com/anurag-dhamala/go-task-cli/db"
	"github.com/spf13/cobra"
)

func GetListCmd() *cobra.Command {
	var list = &cobra.Command{
		Use:   "list",
		Short: "List all the available todo tasks",
		Run: func(cmd *cobra.Command, args []string) {
			db.GetAll()
		},
	}
	return list
}
