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
	LineBreak     *LineBreak     `xml:"w:br,omitempty"`
}

// RunProperties encapsulates visual properties of a run
type RunProperties struct {
	XMLName xml.Name `xml:"w:rPr"`
	Color   *Color   `xml:"w:color,omitempty"`
	Size    *Size    `xml:"w:sz,omitempty"`
	Fonts   *Fonts   `xml:"w:rFonts,omitempty"`
}

func (r *Run) checkPropsInit() {
	if r.RunProperties == nil {
		r.RunProperties = &RunProperties{}
	}
}

// Color allows to set run color
func (r *Run) Color(color string) *Run {
	r.checkPropsInit()
	r.RunProperties.Color = &Color{
		Val: color,
	}

	return r
}

// Size allows to set run size
func (r *Run) Size(size int) *Run {
	r.checkPropsInit()
	r.RunProperties.Size = &Size{
		Val: size * 2,
	}
	return r
}

// Font allows to set run font
func (r *Run) fonts(fontName string, fontHint string) *Run {
	r.checkPropsInit()
	r.RunProperties.Fonts = &Fonts{
		Hint:     fontHint,
		Ascii:    fontName,
		Cs:       fontName,
		EastAsia: fontName,
		HAnsi:    fontName,
	}
	return r
}
func (r *Run) FontDefault(fontName string) *Run {
	r.fonts(fontName, "default")
	return r
}
func (r *Run) FontEastAsia(fontName string) *Run {
	r.fonts(fontName, "eastAsia")
	return r
}
func (r *Run) FontCS(fontName string) *Run {
	r.fonts(fontName, "cs")
	return r
}

func (r *Run) SetPageBreak() *Run {
	r.LineBreak = &LineBreak{
		Type: "page",
	}
	return r
}
