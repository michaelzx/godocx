package section

import "encoding/xml"

const OrientLandscape Orient = "landscape" // 横向
const OrientPortrait Orient = "portrait"   // 纵向
const A4Width int = 11906
const A4Height int = 16838

type PgSz struct {
	XMLName xml.Name `xml:"w:pgSz"`
	Height  *int     `xml:"w:h,attr"`
	Width   *int     `xml:"w:w,attr"`
	Orient  *Orient  `xml:"w:orient,attr"`
}

func (sz *PgSz) SetWidth(w int) *PgSz {
	sz.Width = &w
	return sz
}
func (sz *PgSz) SetHeight(h int) *PgSz {
	sz.Height = &h
	return sz
}

func (sz *PgSz) SetA4Size() *PgSz {
	w := A4Width
	h := A4Height
	o := OrientPortrait
	sz.Width = &w
	sz.Height = &h
	sz.Orient = &o
	return sz
}
func (sz *PgSz) SetOrient(val Orient) *PgSz {
	sz.Orient = &val
	return sz
}
