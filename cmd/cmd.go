package cmd

import (
	"github.com/pettermachado/slidegen/lib/errcheck"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slidegen",
	Short: "slidegen is a Gnome desktop slidehow generator and manager",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			errcheck.Exit(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errcheck.Exit(err)
	}
}
