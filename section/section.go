package section

import "encoding/xml"

type Orient string

type Section struct {
	XMLName xml.Name `xml:"w:sectPr"`
	WPgSz   *PgSz    `xml:"w:pgSz,omitempty"`
}

func (s *Section) PgSz() *PgSz {
	if s.WPgSz == nil {
		s.WPgSz = &PgSz{}
	}
	return s.WPgSz
}
