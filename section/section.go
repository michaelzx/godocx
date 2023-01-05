package section

import "encoding/xml"

type Orient string

const OrientLandscape Orient = "landscape"
const OrientPortrait Orient = "portrait"
const A4Width int = 11906
const A4Height int = 16838

type Section struct {
	XMLName xml.Name `xml:"w:sectPr"`
	PgSz    *PgSz
}
type PgSz struct {
	XMLName xml.Name `xml:"w:pgSz"`
	H       *int     `xml:"w:h,attr"`
	W       *int     `xml:"w:w,attr"`
	Orient  *Orient  `xml:"w:orient,attr"`
}

func (s *Section) checkPgSzInit() {
	if s.PgSz == nil {
		s.PgSz = &PgSz{}
	}
}
func (s *Section) Width(w int) *Section {
	s.checkPgSzInit()
	s.PgSz.W = &w
	return s
}
func (s *Section) Height(h int) *Section {
	s.checkPgSzInit()
	s.PgSz.H = &h
	return s
}

func (s *Section) A4Size() *Section {
	s.checkPgSzInit()
	w := A4Width
	h := A4Height
	s.PgSz.W = &w
	s.PgSz.H = &h
	return s
}
func (s *Section) Orient(val Orient) *Section {
	s.checkPgSzInit()
	s.PgSz.Orient = &val
	return s
}
