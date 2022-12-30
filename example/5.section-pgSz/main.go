package main

import (
	"github.com/michaelzx/godocx"
	"os"
)

func main() {
	docx := godocx.New()
	// add text
	s := docx.AddSection()
	// A4-Landscape
	s.Width(16838).Height(11906).Orient(godocx.OrientLandscape)
	f, err := os.Create("./tmp/5.zip")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = docx.Write(f)
	if err != nil {
		panic(err)
	}
}
