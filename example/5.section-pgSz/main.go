package main

import (
	"github.com/michaelzx/godocx"
	"github.com/michaelzx/godocx/section"
	"os"
)

func main() {
	docx := godocx.New()
	body := docx.Body()
	p0 := body.AddParagraph()
	p0.AddText("test1")
	p1 := body.AddParagraph()
	p1.AddSection().Width(section.A4Height).Height(section.A4Width).Orient(section.OrientLandscape)
	p1.AddText("test2")

	f, err := os.Create("./tmp/5.docx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = docx.Write(f)
	if err != nil {
		panic(err)
	}
}
