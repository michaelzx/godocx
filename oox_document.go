package godocx

import "encoding/xml"

type Document struct {
	XMLName xml.Name `xml:"w:document"`
	XMLW    string   `xml:"xmlns:w,attr"`
	XMLR    string   `xml:"xmlns:r,attr"`
	Body    *Body
}
type Body struct {
	XMLName  xml.Name `xml:"w:body"`
	Children []any
}
