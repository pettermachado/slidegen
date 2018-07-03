package backgrounds

import (
	"encoding/xml"
	"time"
)

// Transition is a transition between two Static sections in the slideshow
type Transition struct {
	XMLName  xml.Name    `xml:"transition"`
	Type     string      `xml:"type,attr"`
	Duration PrettyFloat `xml:"duration"`
	From     string      `xml:"from"`
	To       string      `xml:"to"`
}

func NewTransition(duration time.Duration, from, to string) Transition {
	return Transition{
		Type:     "overlay",
		Duration: PrettyFloat(duration.Seconds()),
		From:     from,
		To:       to,
	}
}
