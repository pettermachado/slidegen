package backgrounds

import (
	"encoding/xml"
	"fmt"
)

// PrettyFloat is a float64 that will MarshalXML with one decimal point
type PrettyFloat float64

func (f PrettyFloat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(fmt.Sprintf("%.1f", f), start)
}
