package paragraph

import (
	"encoding/xml"
)

// http://www.officeopenxml.com/WPtextFormatting.php

// A Run is part of a paragraph that has its own style. It could be
// a piece of text in bold, or a link
type Run struct {
	XMLName       xml.Name       `xml:"w:r"`
	RunProperties *RunProperties `xml:"w:rPr,omitempty"`
	Text          *Text          `xml:"w:t,omitempty"`
}

// RunProperties encapsulates visual properties of a run
type RunProperties struct {
	XMLName xml.Name `xml:"w:rPr"`
	Color   *Color   `xml:"w:color,omitempty"`
	Size    *Size    `xml:"w:sz,omitempty"`
	Fonts   *Fonts   `xml:"w:rFonts,omitempty"`
}

// Fonts contains the font family
// http://officeopenxml.com/WPtextFonts.php
// http://www.datypic.com/sc/ooxml/a-w_hint-1.html
type Fonts struct {
	XMLName  xml.Name `xml:"w:rFonts"`
	Hint     string   `xml:"w:hint,attr,omitempty"` // default=High ANSI Font,eastAsia=East Asian Font,cs=Complex Script Font
	Ascii    string   `xml:"w:ascii,attr,omitempty"`
	Cs       string   `xml:"w:cs,attr,omitempty"`
	EastAsia string   `xml:"w:eastAsia,attr,omitempty"`
	HAnsi    string   `xml:"w:hAnsi,attr,omitempty"`
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

// Font allows to set run font
func (r *Run) Font(fontName string) *Run {
	r.RunProperties.Fonts = &Fonts{
		Hint:     "default",
		Ascii:    fontName,
		Cs:       fontName,
		EastAsia: fontName,
		HAnsi:    fontName,
	}
	return r
}
func (r *Run) FontEastAsia(fontName string) *Run {
	r.RunProperties.Fonts = &Fonts{
		Hint:     "eastAsia",
		Ascii:    fontName,
		Cs:       fontName,
		EastAsia: fontName,
		HAnsi:    fontName,
	}
	return r
}
func (r *Run) FontCS(fontName string) *Run {
	r.RunProperties.Fonts = &Fonts{
		Hint:     "cs",
		Ascii:    fontName,
		Cs:       fontName,
		EastAsia: fontName,
		HAnsi:    fontName,
	}
	return r
}
