package oox

import "encoding/xml"

type Relationships struct {
	XMLName  xml.Name        `xml:"Relationships"`
	Xmlns    string          `xml:"xmlns,attr"`
	Children []*Relationship `xml:"Relationship"`
}

type Relationship struct {
	XMLName    xml.Name `xml:"Relationship"`
	ID         string   `xml:"Id,attr"`
	Type       string   `xml:"Type,attr"`
	Target     string   `xml:"Target,attr"`
	TargetMode string   `xml:"TargetMode,attr,omitempty"`
}
