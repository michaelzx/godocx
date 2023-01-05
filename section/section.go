package section

import "encoding/xml"

type Orient string

type Section struct {
	XMLName xml.Name `xml:"w:sectPr"`
	PgSz    *PgSz    `xml:"w:pgSz,omitempty"`
}

func (s *Section) AddPgSz() *PgSz {
	if s.PgSz == nil {
		s.PgSz = &PgSz{}
	}
	return s.PgSz
}
