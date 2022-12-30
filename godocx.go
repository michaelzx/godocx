package godocx

import (
	"encoding/xml"
)

func New() *Docx {
	docx := &Docx{
		Document: Document{
			XMLName: xml.Name{
				Space: "w",
			},
			XMLW: XMLNS_W,
			XMLR: XMLNS_R,
			Body: &Body{
				XMLName: xml.Name{
					Space: "w",
				},
				Paragraphs: make([]*Paragraph, 0),
			},
		},
		Relationships: Relationships{
			Xmlns: XMLNS,
			Children: []*Relationship{
				{
					ID:     "rId1",
					Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles`,
					Target: "styles.xml",
				},
				{
					ID:     "rId2",
					Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme`,
					Target: "theme/theme1.xml",
				},
				{
					ID:     "rId3",
					Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable`,
					Target: "fontTable.xml",
				},
			},
		},
		NextRelationshipId: 4,
	}
	return docx
}
