package cmd

import (
	"fmt"

	"github.com/pettermachado/slidegen/lib/errcheck"
	"github.com/pettermachado/slidegen/lib/wallpapers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the backgrounds installed by this tool",
	Run: func(cmd *cobra.Command, args []string) {
		ws, err := wallpapers.Load()
		if err != nil {
			errcheck.Exit(err)
		}
		for _, w := range ws.W {
			fmt.Printf("%s\n", w.Name)
		}
	},
}
