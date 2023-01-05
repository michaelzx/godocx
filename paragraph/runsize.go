package paragraph

import "encoding/xml"

// Size contains the font size
type Size struct {
	XMLName xml.Name `xml:"w:sz"`
	Val     int      `xml:"w:val,attr"`
}
