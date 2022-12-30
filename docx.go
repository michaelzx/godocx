package godocx

import (
	"archive/zip"
	"encoding/xml"
	"github.com/michaelzx/godocx/tpl"
	"io"
	"os"
)

type Docx struct {
	Document           Document
	Relationships      Relationships
	NextRelationshipId int
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

// AddParagraph add new paragraph
func (f *Docx) AddParagraph() *Paragraph {
	p := &Paragraph{
		Data: make([]interface{}, 0),
		file: f,
	}

	f.Document.Body.Paragraphs = append(f.Document.Body.Paragraphs, p)
	return p
}

func (f *Docx) pack(zipWriter *zip.Writer) (err error) {
	// ooxml文档地址：http://officeopenxml.com/anatomyofOOXML.php
	files := map[string]string{}
	files["[Content_Types].xml"] = tpl.CONTENT
	files["_rels/.rels"] = tpl.REL
	files["docProps/app.xml"] = tpl.DOCPROPS_APP
	files["docProps/core.xml"] = tpl.DOCPROPS_CORE
	files["word/theme/theme1.xml"] = tpl.WORD_THEME_THEME
	files["word/styles.xml"] = tpl.WORD_STYLE
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
