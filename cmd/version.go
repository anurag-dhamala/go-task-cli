package cmd

import "github.com/spf13/cobra"

func GetVersionCmd() *cobra.Command {
	var version = &cobra.Command{
		Use:   "version",
		Short: "get the version of go task cli",
	}
	return version
}
