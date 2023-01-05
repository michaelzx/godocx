package paragraph

import (
	"encoding/xml"
)

// The Text object contains the actual text
type Text struct {
	XMLName  xml.Name `xml:"w:t"`
	XMLSpace string   `xml:"xml:space,attr,omitempty"`
	Text     string   `xml:",chardata"`
}
