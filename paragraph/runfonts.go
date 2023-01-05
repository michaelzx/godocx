package paragraph

import "encoding/xml"

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
