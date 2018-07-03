package backgrounds

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/pettermachado/slidegen/lib/errcheck"
)

const (
	backgroundsDir = "/usr/share/backgrounds/slidegen"
)

type Background struct {
	Name       string
	Images     []string
	Duration   time.Duration
	Transition time.Duration
	StartTime  time.Time
}

// New returns a new Background
func New(name string, images []string, duration, transiton time.Duration) *Background {
	return &Background{
		Name:       name,
		Images:     images,
		Duration:   duration,
		Transition: transiton,
		StartTime:  time.Now(),
	}
}

func (bg Background) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	name := xml.Name{Local: "background"}
	if err := e.EncodeToken(xml.StartElement{Name: name}); err != nil {
		return err
	}

	st := NewStartTime(bg.StartTime)
	if err := e.Encode(st); err != nil {
		return err
	}

	for i := range bg.Images {
		from := bg.Images[i]
		to := bg.Images[(i+1)%len(bg.Images)]

		st := NewStatic(bg.Duration, from)
		if err := e.Encode(st); err != nil {
			return err
		}
		tr := NewTransition(bg.Transition, from, to)
		if err := e.Encode(tr); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: name})
}

// Store stores a Background and returns the full path to the
// saved backgrounds file and nil, or an empty string and an error.
func Store(bg *Background) (string, error) {
	dir := filepath.Join(backgroundsDir, bg.Name)

	// Make sure directory hierarchy exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// Empty directory and recreate it
	if err := os.RemoveAll(dir); err != nil {
		return "", err
	}
	if err := os.Mkdir(dir, 0755); err != nil {
		return "", err
	}

	// Copy images to make them accessable for all users
	// and the lock screen
	for i, image := range bg.Images {
		s, err := os.Open(image)
		if err != nil {
			return "", err
		}
		defer errcheck.Close(s)

		target := path.Join(dir, filepath.Base(image))
		t, err := os.Create(target)
		if err != nil {
			return "", err
		}
		defer errcheck.Close(t)

		if _, err := io.Copy(t, s); err != nil {
			return "", err
		}
		// Images must be readable by all
		if err := os.Chmod(target, 0644); err != nil {
			return "", err
		}

		bg.Images[i] = target
	}

	b, err := xml.MarshalIndent(bg, "", "  ")
	if err != nil {
		return "", err
	}
	b = []byte(xml.Header + string(b))

	filename := filepath.Join(dir, "background.xml")
	if err = ioutil.WriteFile(filename, b, 0644); err != nil {
		return "", err
	}

	return filename, nil
}
