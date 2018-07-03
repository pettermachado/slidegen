package wallpapers

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const (
	WallpaperDoctype = `<!DOCTYPE wallpapers SYSTEM "gnome-wp-list.dtd">` + "\n"

	wallpapersFile = "/usr/share/gnome-background-properties/slidegen.xml"
)

type Wallpaper struct {
	XMLName   xml.Name `xml:"wallpaper"`
	Name      string   `xml:"name"`
	Filename  string   `xml:"filename"`
	Options   string   `xml:"options"`
	Pcolor    string   `xml:"pcolor"`
	Scolor    string   `xml:"scolor"`
	ShadeType string   `xml:"shade_type"`
}

func New(name, filename string) Wallpaper {
	return Wallpaper{
		Name:      name,
		Filename:  filename,
		Options:   "zoom",
		Pcolor:    "#ffffff",
		Scolor:    "#000000",
		ShadeType: "solid",
	}
}

func (w Wallpaper) Valid() bool {
	return w.Name != "" && w.Filename != ""
}

type Wallpapers struct {
	XMLName xml.Name    `xml:"wallpapers"`
	W       []Wallpaper `xml:"wallpaper"`
}

func (ws *Wallpapers) Add(w Wallpaper) {
	for i := range ws.W {
		if ws.W[i].Name == w.Name {
			ws.W[i] = w
			return
		}
	}
	ws.W = append(ws.W, w)
}

func (ws *Wallpapers) Get(name string) Wallpaper {
	for _, w := range ws.W {
		if w.Name == name {
			return w
		}
	}
	return Wallpaper{}
}

func (ws *Wallpapers) Remove(name string) {
	var _ws []Wallpaper
	for _, w := range ws.W {
		if w.Name == name {
			continue
		}
		_ws = append(_ws, w)
	}
	ws.W = _ws
}

func (ws *Wallpapers) Sort() {
	sort.Slice(ws.W, func(i, j int) bool {
		return strings.ToLower(ws.W[i].Name) < strings.ToLower(ws.W[j].Name)
	})
}

func Load() (*Wallpapers, error) {
	_, err := os.Stat(wallpapersFile)
	if os.IsNotExist(err) {
		return &Wallpapers{}, nil
	}
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(wallpapersFile)
	if err != nil {
		return nil, err
	}
	var ws Wallpapers
	return &ws, xml.Unmarshal(b, &ws)
}

func Store(ws *Wallpapers) error {
	ws.Sort()

	b, err := xml.MarshalIndent(ws, "", "\t")
	if err != nil {
		return err
	}
	b = []byte(xml.Header + WallpaperDoctype + string(b))
	return ioutil.WriteFile(wallpapersFile, b, 0644)
}
