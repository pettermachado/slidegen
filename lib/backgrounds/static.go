package backgrounds

import (
	"encoding/xml"
	"time"
)

// Static is a slide in the slideshow, it has an image file and a duration
type Static struct {
	XMLName  xml.Name    `xml:"static"`
	Duration PrettyFloat `xml:"duration"`
	File     string      `xml:"file"`
}

func NewStatic(duration time.Duration, file string) Static {
	return Static{
		Duration: PrettyFloat(duration.Seconds()),
		File:     file,
	}
}
