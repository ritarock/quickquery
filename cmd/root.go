package cmd

import (
	"os"

	"github.com/ritarock/quickquery/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quickquery",
	Short: "quickquery can search from csv like sql",
	RunE: func(cmd *cobra.Command, args []string) error {
		handler := cli.NewHandler()
		if err := handler.Run(args); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
