package cmd

import (
	"os"
	"path/filepath"

	"github.com/pettermachado/slidegen/lib/errcheck"
	"github.com/pettermachado/slidegen/lib/wallpapers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:     "remove [backgrounds]",
	Aliases: []string{"rm"},
	Short:   "Remove backgrounds",
	Run: func(cmd *cobra.Command, args []string) {
		ws, err := wallpapers.Load()
		if err != nil {
			errcheck.Exit(err)
		}
		for _, n := range args {
			w := ws.Get(n)
			if !w.Valid() {
				continue
			}
			ws.Remove(n)
			if err := os.RemoveAll(filepath.Dir(w.Filename)); err != nil {
				errcheck.Exit(err)
			}
		}
		if err := wallpapers.Store(ws); err != nil {
			errcheck.Exit(err)
		}
	},
}
