package godocx

import "encoding/xml"

type Paragraph struct {
	XMLName xml.Name      `xml:"w:p"`
	Data    []interface{} // todo：可否用泛型替代？
	file    *Docx
}

// AddText add text to paragraph
func (p *Paragraph) AddText(text string) *Run {
	t := &Text{
		Text: text,
	}
	run := &Run{
		Text:          t,
		RunProperties: &RunProperties{},
	}
	p.Data = append(p.Data, run)

	return run
}
