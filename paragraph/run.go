package paragraph

import (
	"encoding/xml"
)

// A Run is part of a paragraph that has its own style. It could be
// a piece of text in bold, or a link
type Run struct {
	XMLName       xml.Name `xml:"w:r"`
	RunProperties *RunProperties
	InstrText     string `xml:"w:instrText,omitempty"`
	Text          *Text
}

// RunProperties encapsulates visual properties of a run
type RunProperties struct {
	XMLName  xml.Name  `xml:"w:r,omitempty"`
	Color    *Color    `xml:"w:color,omitempty"`
	Size     *Size     `xml:"w:sz,omitempty"`
	RunStyle *RunStyle `xml:"w:rStyle,omitempty"`
	Style    *Style    `xml:"w:pStyle,omitempty"`
}

// RunStyle contains styling for a run
type RunStyle struct {
	XMLName xml.Name `xml:"w:rStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Style contains styling for a paragraph
type Style struct {
	XMLName xml.Name `xml:"w:pStyle,omitempty"`
	Val     string   `xml:"w:val,attr"`
}

// Color contains the sound of music. :D
// I'm kidding. It contains the color
type Color struct {
	XMLName xml.Name `xml:"w:color"`
	Val     string   `xml:"w:val,attr"`
}

// Size contains the font size
type Size struct {
	XMLName xml.Name `xml:"w:sz"`
	Val     int      `xml:"w:val,attr"`
}

// Color allows to set run color
func (r *Run) Color(color string) *Run {
	r.RunProperties.Color = &Color{
		Val: color,
	}

	return r
}

// Size allows to set run size
func (r *Run) Size(size int) *Run {
	r.RunProperties.Size = &Size{
		Val: size * 2,
	}
	return r
}
