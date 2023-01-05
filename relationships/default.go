package relationships

const XMLNS = `http://schemas.openxmlformats.org/package/2006/relationships`

func Default() *Relationships {
	return &Relationships{
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
		RelationshipId: 3,
	}
}
