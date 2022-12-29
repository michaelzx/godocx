package godocx

import (
	"encoding/xml"
	"github.com/michaelzx/godocx/oox"
)

func New() *oox.Docx {
	docx := &oox.Docx{
		Document: oox.Document{
			XMLName: xml.Name{
				Space: "w",
			},
			XMLW: oox.XMLNS_W,
			XMLR: oox.XMLNS_R,
			Body: &oox.Body{
				XMLName: xml.Name{
					Space: "w",
				},
				Paragraphs: make([]*oox.Paragraph, 0),
			},
		},
		Relationships: oox.Relationships{
			Xmlns: oox.XMLNS,
			Children: []*oox.Relationship{
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
