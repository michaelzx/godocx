package main

import (
	"github.com/michaelzx/godocx"
	"github.com/michaelzx/godocx/section"
	"os"
)

func main() {
	docx := godocx.New()
	body := docx.Body()

	p1 := body.AddParagraph()
	p1.AddSection().AddPgSz().SetA4Size()
	p1.AddText("page1")

	p2 := body.AddParagraph()
	p2.AddSection().AddPgSz().SetWidth(section.A4Height).SetHeight(section.A4Width).SetOrient(section.OrientLandscape)
	p2.AddText("page2")

	p3 := body.AddParagraph()
	p3.AddSection().AddPgSz().SetA4Size()
	p3.AddText("page3")

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
