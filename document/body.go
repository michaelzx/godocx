package document

import (
	"encoding/xml"
	"github.com/michaelzx/godocx/paragraph"
	"github.com/michaelzx/godocx/relationships"
	"github.com/michaelzx/godocx/section"
)

type Body struct {
	XMLName       xml.Name                     `xml:"w:body"`
	Relationships *relationships.Relationships `xml:"-"`
	Children      []any
}

// AddParagraph add new paragraph
func (b *Body) AddParagraph() *paragraph.Paragraph {
	p := paragraph.New(b.Relationships)
	b.Children = append(b.Children, p)
	return p
}

// Section set body section config
func (b *Body) Section() *section.Section {
	s := &section.Section{}
	b.Children = append(b.Children, s)
	return s
}
