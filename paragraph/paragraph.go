package paragraph

import (
	"encoding/xml"
	"github.com/michaelzx/godocx/relationships"
	"github.com/michaelzx/godocx/section"
)

// http://www.officeopenxml.com/WPparagraph.php

type Paragraph struct {
	XMLName       xml.Name                     `xml:"w:p"`
	Relationships *relationships.Relationships `xml:"-"`
	Properties    Properties                   `xml:"w:pPr,omitempty"`
	Children      []any
}
type Properties struct {
	XMLName  xml.Name `xml:"w:pPr"`
	Children []any
}

func New(r *relationships.Relationships) *Paragraph {
	p := &Paragraph{
		Relationships: r,
		Children:      make([]interface{}, 0),
		Properties: Properties{
			Children: make([]any, 0, 0),
		},
	}
	return p
}

// AddText add text to paragraph
func (p *Paragraph) AddText(text string) *Run {
	t := &Text{
		Text: text,
	}
	run := &Run{
		Text: t,
	}
	p.Children = append(p.Children, run)

	return run
}

// AddSection add section to paragraph
func (p *Paragraph) AddSection() *section.Section {
	s := &section.Section{}
	p.Properties.Children = append(p.Properties.Children, s)
	return s
}
