package cmd

import (
	"os"
	"quickquery/interface/cli"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quickquery",
	Short: "quickquery can search from csv like sql",
	RunE: func(cmd *cobra.Command, args []string) error {
		handler := cli.NewHandler()
		return handler.Run(args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
