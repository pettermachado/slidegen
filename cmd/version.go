package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of slidegen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("slidenen v0.9")
	},
}
