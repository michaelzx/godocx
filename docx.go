package godocx

import (
	"archive/zip"
	"encoding/xml"
	"github.com/michaelzx/godocx/document"
	"github.com/michaelzx/godocx/relationships"
	"github.com/michaelzx/godocx/tpl"
	"io"
	"os"
)

type Docx struct {
	Document      document.Document
	Relationships *relationships.Relationships
}

const (
	XmlNsW = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XmlNsR = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
)

func New() *Docx {
	docx := &Docx{
		Document: document.Document{
			XMLName: xml.Name{
				Space: "w",
			},
			XmlNsW: XmlNsW,
			XmlNsR: XmlNsR,
			Body: &document.Body{
				XMLName: xml.Name{
					Space: "w",
				},
				Children: make([]any, 0),
			},
		},
		Relationships: relationships.Default(),
	}
	return docx
}

// Body get body
func (f *Docx) Body() (err *document.Body) {
	return f.Document.Body
}

// Save save file to path
func (f *Docx) Save(path string) (err error) {
	fzip, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

// Write allows to save a docx to a writer
func (f *Docx) Write(writer io.Writer) (err error) {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

func (f *Docx) pack(zipWriter *zip.Writer) (err error) {
	// ooxml文档地址：http://officeopenxml.com/anatomyofOOXML.php
	files := map[string]string{}
	files["[Content_Types].xml"] = tpl.ContentTypes()
	files["_rels/.rels"] = tpl.Rels()
	files["docProps/app.xml"] = tpl.DocPropsApp()
	files["docProps/core.xml"] = tpl.DocPropsCore()
	files["word/theme/theme1.xml"] = tpl.WordThemeTheme1()
	files["word/styles.xml"] = tpl.WordStyles()
	files["word/_rels/document.xml.rels"], err = marshal(f.Relationships)
	if err != nil {
		return err
	}
	files["word/document.xml"], err = marshal(f.Document)
	if err != nil {
		return err
	}

	for path, data := range files {
		w, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(data))
		if err != nil {
			return err
		}
	}

	return
}

func marshal(data interface{}) (out string, err error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return
	}
	out = xml.Header + string(body)
	return
}
