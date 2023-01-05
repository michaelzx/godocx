package relationships

import (
	"encoding/xml"
	"fmt"
)

type Relationships struct {
	XMLName        xml.Name        `xml:"Relationships"`
	Xmlns          string          `xml:"xmlns,attr"`
	Children       []*Relationship `xml:"Relationship"`
	RelationshipId int             `xml:"-"`
}

type Relationship struct {
	XMLName    xml.Name `xml:"Relationship"`
	ID         string   `xml:"Id,attr"`
	Type       string   `xml:"Type,attr"`
	Target     string   `xml:"Target,attr"`
	TargetMode string   `xml:"TargetMode,attr,omitempty"`
}

func (r *Relationships) NextRelationshipIDStr() string {
	r.RelationshipId++
	return fmt.Sprintf("rId%d", r.RelationshipId)
}
