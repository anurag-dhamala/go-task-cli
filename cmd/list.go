package cmd

import "github.com/spf13/cobra"

func GetListCmd() *cobra.Command {
	var list = &cobra.Command{
		Use:   "list",
		Short: "List all the available todo tasks",
	}
	return list
}
