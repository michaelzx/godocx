package document

import "encoding/xml"

type Document struct {
	XMLName xml.Name `xml:"w:document"`
	XmlNsW  string   `xml:"xmlns:w,attr"`
	XmlNsR  string   `xml:"xmlns:r,attr"`
	Body    *Body
}
