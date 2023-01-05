package paragraph

import "encoding/xml"

// Color contains the sound of music. :D
// I'm kidding. It contains the color
type Color struct {
	XMLName xml.Name `xml:"w:color"`
	Val     string   `xml:"w:val,attr"`
}
