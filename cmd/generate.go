package cmd

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/pettermachado/slidegen/lib/backgrounds"
	"github.com/pettermachado/slidegen/lib/errcheck"
	"github.com/pettermachado/slidegen/lib/wallpapers"
	"github.com/spf13/cobra"
)

var name string
var duration, transition int

func init() {
	generateCmd.Flags().StringVar(&name, "name", "", "the slideshow name. Using the folder name as default.")
	generateCmd.Flags().IntVar(&duration, "duration", 60*30, "slide duration in seconds")
	generateCmd.Flags().IntVar(&transition, "transition", 2, "transition duration seconds")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate folder",
	Short: "Generate a new slideshow background",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		folder := args[0]
		files, err := ioutil.ReadDir(folder)
		if err != nil {
			errcheck.Exit(err)
		}
		images := ImagePaths(folder, files)

		path, err := filepath.Abs(folder)
		if err != nil {
			errcheck.Exit(err)
		}
		if name == "" {
			name = filepath.Base(path)
		}

		durationD := time.Duration(duration) * time.Second
		transitionD := time.Duration(transition) * time.Second

		fmt.Printf("Generating\n\tName: %s\n\tImage(s): %d from %s\n\tDuration/Transition: %s/%s\n", name, len(images), path, durationD, transitionD)

		bg := backgrounds.New(name, images, durationD, transitionD)
		file, err := backgrounds.Store(bg)
		if err != nil {
			errcheck.Exit(err)
		}

		ws, err := wallpapers.Load()
		if err != nil {
			errcheck.Exit(err)
		}
		ws.Add(wallpapers.New(name, file))
		if err := wallpapers.Store(ws); err != nil {
			errcheck.Exit(err)
		}
	},
}

func ImagePaths(folder string, files []os.FileInfo) []string {
	var out []string
	for _, file := range files {
		if !file.Mode().IsRegular() {
			continue
		}
		p, err := filepath.Abs(path.Join(folder, file.Name()))
		if err != nil {
			fmt.Printf("skipping %q: %s\n", file.Name(), err)
			continue
		}
		r, err := os.Open(p)
		if err != nil {
			fmt.Printf("skipping %q: %s\n", file.Name(), err)
			continue
		}
		if _, _, err := image.DecodeConfig(r); err != nil {
			fmt.Printf("skipping %q: %s\n", file.Name(), err)
			continue
		}
		out = append(out, p)
	}
	return out
}
