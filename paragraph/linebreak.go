package paragraph

import "encoding/xml"

// http://www.officeopenxml.com/WPtextSpecialContent-break.php

// LineBreak  A line break is specified with the <w:br> element.
// It specifies a break which overrides the normal line breaking.
// (Normal line breaking occurs after a breaking space or optional hyphen character.)
type LineBreak struct {
	XMLName xml.Name `xml:"w:br"`
	Type    string   `xml:"w:type,attr,omitempty"`  // column,page,textWrapping
	Clear   string   `xml:"w:clear,attr,omitempty"` // all,left,none,right
}
