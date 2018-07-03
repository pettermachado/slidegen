package backgrounds

import (
	"encoding/xml"
	"time"
)

type StartTime struct {
	XMLName xml.Name `xml:"starttime"`
	Hour    int      `xml:"hour"`
	Minute  int      `xml:"minute"`
	Second  int      `xml:"second"`
}

func NewStartTime(t time.Time) StartTime {
	return StartTime{
		Hour:   t.Hour(),
		Minute: t.Minute(),
		Second: t.Second(),
	}
}
