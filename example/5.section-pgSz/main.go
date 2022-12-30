package main

import (
	"github.com/michaelzx/godocx"
	"os"
)

func main() {
	docx := godocx.New()
	docx.AddSection().Width(16838).Height(11906).Orient(godocx.OrientLandscape)

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
