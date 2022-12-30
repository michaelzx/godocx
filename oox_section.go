package godocx

import "encoding/xml"

type Orient string

const OrientLandscape Orient = "landscape"
const OrientPortrait Orient = "portrait"

type Section struct {
	XMLName xml.Name `xml:"w:sectPr"`
	PgSz    *SectionPgSz
}
type SectionPgSz struct {
	XMLName xml.Name `xml:"w:pgSz"`
	H       *int     `xml:"w:h,attr"`
	W       *int     `xml:"w:w,attr"`
	Orient  *Orient  `xml:"w:orient,attr"`
}

func (s *Section) checkPgSzInit() {
	if s.PgSz == nil {
		s.PgSz = &SectionPgSz{}
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
func (s *Section) Orient(val Orient) *Section {
	s.checkPgSzInit()
	s.PgSz.Orient = &val
	return s
}
